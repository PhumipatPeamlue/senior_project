package drug_doc_service

import (
	"admin_management_service/errs"
	"admin_management_service/models"
	"admin_management_service/ports"
	"errors"
	"fmt"
	"github.com/elastic/go-elasticsearch/v8/typedapi/types"
	"log"
	"net/http"
	"os"
	"time"
)

type drugDocService struct {
	drugDocRepo  ports.DrugDocRepo
	docImageRepo ports.DocImageRepo
}

func (r *drugDocService) handleEsErr(err error) (newErr error) {
	var esErr types.ElasticsearchError
	if errors.As(err, &esErr) {
		switch esErr.Status {
		case http.StatusNotFound:
			newErr = errs.DrugDocNotFoundError
		default:
			newErr = errs.UnexpectedError
		}
	} else {
		newErr = errs.UnexpectedError
	}
	return
}

func (r *drugDocService) handleMysqlErr(err error) (newErr error) {
	if errors.Is(err, errs.DocImageNotFoundError) {
		newErr = errs.DocImageNotFoundError
	} else {
		newErr = errs.UnexpectedError
	}
	return
}

func (r *drugDocService) CreateDrugDoc(doc models.DrugDoc, imageName string, imagePath string) (err error) {
	now := time.Now()
	doc.CreateAt = &now
	doc.UpdateAt = &now
	docID, err := r.drugDocRepo.Index(doc)
	if err != nil {
		log.Println(err)
		err = r.handleEsErr(err)
		return
	}

	docImage := models.DocImage{
		DocID: docID,
		Name:  imageName,
		Path:  imagePath,
	}
	if err = r.docImageRepo.Insert(docImage); err != nil {
		log.Println(err)
		err = r.handleMysqlErr(err)
	}

	return
}

func (r *drugDocService) GetDrugDoc(docID string) (doc models.DrugDocWithID, imageURL string, err error) {
	doc, err = r.drugDocRepo.Get(docID)
	if err != nil {
		log.Println(err)
		err = r.handleEsErr(err)
		return
	}

	docImage, err := r.docImageRepo.SelectByDocID(docID)
	if err != nil {
		err = r.handleMysqlErr(err)
		return
	}

	if len(docImage.Name) == 0 {
		imageURL = ""
		return
	}
	imageURL = fmt.Sprintf("http://localhost:8080/image/%s", docImage.Name)
	return
}

func (r *drugDocService) SearchDrugDoc(from int, size int, keyword string) (docs []models.DrugDocWithID, total int, err error) {
	if len(keyword) == 0 {
		docs, total, err = r.drugDocRepo.SearchMatchAll(from, size)
	} else {
		docs, total, err = r.drugDocRepo.SearchMatchKeyword(from, size, keyword)
	}

	return
}

func (r *drugDocService) UpdateDrugDoc(doc models.DrugDocWithID, imageName string, imagePath string) (err error) {
	oldDoc, err := r.drugDocRepo.Get(doc.ID)
	if err != nil {
		log.Println(err)
		err = r.handleEsErr(err)
		return
	}

	now := time.Now()
	doc.CreateAt = oldDoc.CreateAt
	doc.UpdateAt = &now
	if err = r.drugDocRepo.Update(doc); err != nil {
		log.Println(err)
		err = r.handleEsErr(err)
		return
	}

	oldDocImage, err := r.docImageRepo.SelectByDocID(doc.ID)
	if err != nil {
		log.Println(err)
		err = r.handleMysqlErr(err)
		return
	}

	newDocImage := models.DocImage{
		DocID: doc.ID,
		Name:  imageName,
		Path:  imagePath,
	}
	if err = r.docImageRepo.Update(newDocImage); err != nil {
		log.Println(err)
		err = r.handleMysqlErr(err)
		return
	}

	if len(oldDocImage.Path) != 0 {
		if err = os.Remove(oldDocImage.Path); err != nil {
			log.Println(err)
			err = errs.UnexpectedError
		}
	}

	return
}

func (r *drugDocService) DeleteDrugDoc(docID string) (err error) {
	if err = r.drugDocRepo.Delete(docID); err != nil {
		log.Println(err)
		err = r.handleEsErr(err)
		return
	}

	docImage, err := r.docImageRepo.SelectByDocID(docID)
	if err != nil {
		log.Println(err)
		err = r.handleMysqlErr(err)
		return
	}

	if err = r.docImageRepo.Delete(docID); err != nil {
		log.Println(err)
		err = r.handleMysqlErr(err)
		return
	}

	if len(docImage.Path) != 0 {
		if err = os.Remove(docImage.Path); err != nil {
			log.Println(err)
			err = errs.UnexpectedError
		}
	}

	return
}

func NewDrugDocService(drugDocRepo ports.DrugDocRepo, docImageRepo ports.DocImageRepo) ports.DrugDocService {
	return &drugDocService{
		drugDocRepo:  drugDocRepo,
		docImageRepo: docImageRepo,
	}
}

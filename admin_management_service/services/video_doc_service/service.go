package video_doc_service

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

type videoDocService struct {
	videoDocRepo ports.VideoDocRepo
	docImageRepo ports.DocImageRepo
}

func (r *videoDocService) handleEsErr(err error) (newErr error) {
	var esErr types.ElasticsearchError
	if errors.As(err, &esErr) {
		switch esErr.Status {
		case http.StatusNotFound:
			newErr = errs.VideoDocNotFoundError
		default:
			newErr = errs.UnexpectedError
		}
	} else {
		newErr = errs.UnexpectedError
	}
	return
}

func (r *videoDocService) handleMysqlErr(err error) (newErr error) {
	if errors.Is(err, errs.DocImageNotFoundError) {
		newErr = errs.DocImageNotFoundError
	} else {
		newErr = errs.UnexpectedError
	}
	return
}

func (r *videoDocService) CreateVideoDoc(doc models.VideoDoc, imageName string, imagePath string) (err error) {
	now := time.Now()
	doc.CreateAt = &now
	doc.UpdateAt = &now
	docID, err := r.videoDocRepo.Index(doc)
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

func (r *videoDocService) GetVideoDoc(docID string) (doc models.VideoDocWithId, imageURL string, err error) {
	doc, err = r.videoDocRepo.Get(docID)
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

func (r *videoDocService) SearchVideoDoc(from int, size int, keyword string) (docs []models.VideoDocWithId, total int, err error) {
	if len(keyword) == 0 {
		docs, total, err = r.videoDocRepo.SearchMatchAll(from, size)
	} else {
		docs, total, err = r.videoDocRepo.SearchMatchKeyword(from, size, keyword)
	}

	return
}

func (r *videoDocService) UpdateVideoDoc(doc models.VideoDocWithId, imageName string, imagePath string) (err error) {
	oldDoc, err := r.videoDocRepo.Get(doc.ID)
	if err != nil {
		log.Println(err)
		err = r.handleEsErr(err)
		return
	}

	now := time.Now()
	doc.CreateAt = oldDoc.CreateAt
	doc.UpdateAt = &now
	if err = r.videoDocRepo.Update(doc); err != nil {
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

func (r *videoDocService) DeleteVideoDoc(docID string) (err error) {
	if err = r.videoDocRepo.Delete(docID); err != nil {
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

func NewVideoDocService(videoDocRepo ports.VideoDocRepo, docImageRepo ports.DocImageRepo) ports.VideoDocService {
	return &videoDocService{
		videoDocRepo: videoDocRepo,
		docImageRepo: docImageRepo,
	}
}

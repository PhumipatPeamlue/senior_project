package video_doc_service

import (
	"errors"
	"fmt"
	"log"
	"mime/multipart"
	"senior_project/admin_management_service/repositories/image_storage_repository"
	"senior_project/admin_management_service/repositories/video_doc_repository"
	"time"
)

type service struct {
	videoDocRepo     video_doc_repository.VideoDocRepository
	imageStorageRepo image_storage_repository.ImageStorageRepository
}

func New(videoDocRepo video_doc_repository.VideoDocRepository, imageStorageRepo image_storage_repository.ImageStorageRepository) VideoDocService {
	return &service{
		videoDocRepo:     videoDocRepo,
		imageStorageRepo: imageStorageRepo,
	}
}

func (s *service) handleErr(err error) error {
	log.Println(err)
	if errors.Is(err, video_doc_repository.ErrVideoDocNotFound) {
		return ErrVideoDocNotFound
	}
	return ErrInternalVideoDocService
}

func (s *service) GetVideoDoc(docID string) (getRes *GetResponse, err error) {
	doc, err := s.videoDocRepo.Get(docID)
	if err != nil {
		err = s.handleErr(err)
		return
	}

	var imageURL string
	if len(doc.ImageName) != 0 {
		imageURL = fmt.Sprintf("http://localhost:8080/image/%s", doc.ImageName)
	}
	getRes = &GetResponse{
		Doc: VideoDoc{
			ID:          docID,
			Title:       doc.Title,
			Description: doc.Description,
			VideoURL:    doc.VideoURL,
			CreateAt:    doc.CreateAt,
			UpdateAt:    doc.UpdateAt,
		},
		ImageURL: imageURL,
	}
	return
}

func (s *service) CreateVideoDoc(newDocReq *NewVideoDocRequest) (err error) {
	now := time.Now()
	doc := video_doc_repository.VideoDoc{
		Title:       newDocReq.Title,
		Description: newDocReq.Description,
		VideoURL:    newDocReq.VideoURL,
		CreateAt:    &now,
		UpdateAt:    &now,
	}
	if err = s.videoDocRepo.Create(doc); err != nil {
		err = s.handleErr(err)
	}
	return
}

func (s *service) CreateVideoDocWithImage(newDocReq *NewVideoDocRequest, file *multipart.File, header *multipart.FileHeader) (err error) {
	now := time.Now()
	imageName := fmt.Sprintf("%s_%s", now.String(), header.Filename)
	doc := video_doc_repository.VideoDoc{
		Title:       newDocReq.Title,
		Description: newDocReq.Description,
		VideoURL:    newDocReq.VideoURL,
		ImageName:   imageName,
		CreateAt:    &now,
		UpdateAt:    &now,
	}
	if err = s.videoDocRepo.Create(doc); err != nil {
		err = s.handleErr(err)
		return
	}

	if err = s.imageStorageRepo.Save(file, imageName); err != nil {
		err = s.handleErr(err)
	}
	return
}

func (s *service) SearchVideoDoc(page int, pageSize int, keyword string) (searchRes *SearchResponse, err error) {
	var docs *[]video_doc_repository.VideoDocWithID
	var total int
	from := (page - 1) * pageSize
	if len(keyword) == 0 {
		docs, total, err = s.videoDocRepo.MatchAll(from, pageSize)
	} else {
		docs, total, err = s.videoDocRepo.MatchKeyword(from, pageSize, keyword)
	}

	searchRes = &SearchResponse{
		Data:  []VideoDoc{},
		Total: total,
	}
	for _, doc := range *docs {
		data := VideoDoc{
			ID:          doc.ID,
			Title:       doc.Title,
			Description: doc.Description,
			VideoURL:    doc.VideoURL,
			CreateAt:    doc.CreateAt,
			UpdateAt:    doc.UpdateAt,
		}
		searchRes.Data = append(searchRes.Data, data)
	}
	return
}

func (s *service) UpdateVideoDoc(updateDocReq *UpdateVideoDocRequest) (err error) {
	var doc *video_doc_repository.VideoDocWithID
	now := time.Now()
	doc, err = s.videoDocRepo.Get(updateDocReq.ID)
	if err != nil {
		err = s.handleErr(err)
		return
	}

	doc.Title = updateDocReq.Title
	doc.Description = updateDocReq.Description
	doc.VideoURL = updateDocReq.VideoURL
	doc.UpdateAt = &now
	if updateDocReq.DeleteImage && len(doc.ImageName) != 0 {
		if err = s.imageStorageRepo.Delete(doc.ImageName); err != nil {
			err = s.handleErr(err)
			return
		}
		doc.ImageName = ""
	}
	if err = s.videoDocRepo.Update(doc); err != nil {
		err = s.handleErr(err)
	}
	return
}

func (s *service) UpdateVideoDocWithImage(updateDocReq *UpdateVideoDocRequest, file *multipart.File, header *multipart.FileHeader) (err error) {
	var doc *video_doc_repository.VideoDocWithID
	now := time.Now()
	imageName := fmt.Sprintf("%s_%s", now.String(), header.Filename)
	doc, err = s.videoDocRepo.Get(updateDocReq.ID)
	if err != nil {
		err = s.handleErr(err)
		return
	}
	if len(doc.ImageName) != 0 {
		if err = s.imageStorageRepo.Delete(doc.ImageName); err != nil {
			err = s.handleErr(err)
			return
		}
	}
	if err = s.imageStorageRepo.Save(file, imageName); err != nil {
		err = s.handleErr(err)
		return
	}

	doc.Title = updateDocReq.Title
	doc.Description = updateDocReq.Description
	doc.VideoURL = updateDocReq.VideoURL
	doc.UpdateAt = &now
	doc.ImageName = imageName
	if err = s.videoDocRepo.Update(doc); err != nil {
		err = s.handleErr(err)
	}
	return
}

func (s *service) DeleteVideoDoc(docID string) (err error) {
	doc, err := s.videoDocRepo.Get(docID)
	if err != nil {
		err = s.handleErr(err)
		return
	}
	if len(doc.ImageName) != 0 {
		if err = s.imageStorageRepo.Delete(doc.ImageName); err != nil {
			err = s.handleErr(err)
			return
		}
	}

	if err = s.videoDocRepo.Delete(docID); err != nil {
		err = s.handleErr(err)
	}
	return
}

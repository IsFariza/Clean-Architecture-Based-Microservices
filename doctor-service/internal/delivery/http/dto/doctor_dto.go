package dto

import "github.com/IsFariza/doctor-service/internal/domain"

type DoctorRequest struct {
	FullName       string `json:"full_name" binding:"required,min=3"`
	Specialization string `json:"specialization" binding:"required"`
	Email          string `json:"email" binding:"required,email"`
}

type DoctorResponse struct {
	ID             string `json:"id"`
	FullName       string `json:"full_name"`
	Specialization string `json:"specialization"`
	Email          string `json:"email"`
}

func ToDomain(req DoctorRequest) *domain.Doctor {
	return &domain.Doctor{
		FullName:       req.FullName,
		Specialization: req.Specialization,
		Email:          req.Email,
	}
}

func FromDomain(d *domain.Doctor) DoctorResponse {
	return DoctorResponse{
		ID:             d.ID,
		FullName:       d.FullName,
		Specialization: d.Specialization,
		Email:          d.Email,
	}
}

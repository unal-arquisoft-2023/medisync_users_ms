package staff

import (
	"medysinc_user_ms/controllers"
	"medysinc_user_ms/domain"
)

type CreateStaffRequest struct {
	controllers.UserCreationRequest `json:",inline"`
}

type UpdateStaffRequest struct {
	controllers.UserUpdateRequest `json:",inline"`
}

type StaffResponse struct {
	controllers.UserResponse `json:",inline"`
}

func makeStaffResponse(p domain.Staff) StaffResponse {
	return StaffResponse{
		UserResponse: controllers.UserResponse{
			ID: p.ID,
			Name: controllers.NameDTO{
				FirstName: p.Name.FirstName,
				LastName:  p.Name.LastName,
			},
			Email:            p.Email,
			Phone:            p.Phone,
			Location:         controllers.LocationDTO{Country: p.Location.Country, City: p.Location.City, Address: p.Location.Address},
			DateOfBirth:      controllers.CivilTime(p.DateOfBirth),
			RegistrationDate: p.RegistrationTimeStamp,
			Status:           controllers.UserStatusDTO(p.Status),
			CardID:           p.CardID,
		},
	}

}

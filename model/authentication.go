package model

type (
	// PasswordModel - PasswordModel
	PasswordModel struct {
		OldPassword        string `json:"old_password" validate:"required,min=8"`
		NewPassword        string `json:"new_password" validate:"required,min=8"`
		ConfirmNewPassword string `json:"confirm_new_password" validate:"required,min=8"`

		VerifyCode  string `json:"verify_code" validate:"required"`
		Email       string `bson:"email" json:"email" validate:"required,email,min=2,max=320"`
		PhoneNumber string `json:"phone_number" validate:"required,number,min=10,max=11"`
	}
)

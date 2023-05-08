package database

import (
	"backend/models"
	"errors"
)

// Creates new user
func CreateUser(user *models.User) error {
	res, err := db.Exec("INSERT INTO users (email, password, authLevel) VALUES (?, ?, ?)", user.Email, user.Password, 1)
	if err != nil {
		return err
	}

	id, err := res.LastInsertId()
	if err != nil {
		return err
	}

	user.Id = uint64(id)
	user.AuthLevel = 1
	return nil
}

// gets user
func GetUser(user *models.User, email string) error {
	row := db.QueryRow("SELECT * FROM users WHERE email=?", email)

	if err := row.Scan(&user.Id, &user.Email, &user.Password, &user.AuthLevel); err != nil {
		return err
	}
	return nil
}

// gets all offers
func GetOffers() ([]models.Offer, error) {
	var offers []models.Offer
	rows, err := db.Query("SELECT offers.id, offers.userID, offers.mealPoints, users.email FROM offers JOIN users on (offers.userID=users.id)")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var offer models.Offer
		if err := rows.Scan(&offer.ID, &offer.UserID, &offer.MealPoints, &offer.Email); err != nil {
			return nil, err
		}
		offers = append(offers, offer)
	}

	return offers, nil
}

// gets all new offers
func CreateNewOffer(userID any, mealPointsOffer uint16) error {
	_, err := db.Exec("INSERT INTO offers (userID, mealPoints) VALUES (?, ?)", userID, mealPointsOffer)
	if err != nil {
		return err
	}

	return nil
}

// gets all delete offers
func DeleteOffer(userID any, offerID uint64) error {
	res, err := db.Exec("DELETE FROM offers WHERE id=? AND userID=?", offerID, userID)
	if err != nil {
		return err
	}

	rows, err := res.RowsAffected()
	if err != nil {
		return err
	}
	if rows != 1 {
		return errors.New("You are unable to delete this offer")
	}
	return nil
}

// gets all updated offers
func UpdateOffer(userID any, offerID uint64, mealPointsOffer uint16) error {
	res, err := db.Exec("UPDATE offers SET mealPoints=? WHERE id=? AND userID=?", mealPointsOffer, offerID, userID)
	if err != nil {
		return err
	}

	rows, err := res.RowsAffected()
	if err != nil {
		return err
	}
	if rows != 1 {
		return errors.New("You are unable to update this offer")
	}

	return nil
}

// gets reported users
func ReportUser(reportID uint64, userID any, message string) error {
	_, err := db.Exec("INSERT INTO reports (reportID, userID, message) VALUES (?, ?, ?)", reportID, userID, message)
	if err != nil {
		return err
	}

	return nil
}

// gets reports
func GetReports() ([]models.Report, error) {
	var reports []models.Report

	rows, err := db.Query("SELECT r.id, r.reportID, r.userID, r.message, u1.email, u2.email FROM reports r JOIN users u1 ON r.reportID= u1.id JOIN users u2 ON r.userID = u2.id")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var report models.Report
		if err := rows.Scan(&report.ID, &report.ReportID, &report.UserID, &report.Message, &report.ReportEmail, &report.UserEmail); err != nil {
			return nil, err
		}
		reports = append(reports, report)
	}

	return reports, nil
}

// bans a user
func BanUser(banUserID uint64) error {
	res, err := db.Exec("UPDATE users SET authLevel=3 WHERE id=?", banUserID)
	if err != nil {
		return err
	}

	rows, err := res.RowsAffected()
	if err != nil {
		return err
	}
	if rows != 1 {
		return errors.New("You are unable to update this offer")
	}

	return nil
}

// verifys a new user
func VerifyUser(userID uint64) error {
	res, err := db.Exec("UPDATE users SET authLevel=2 WHERE id=?", userID)
	if err != nil {
		return err
	}

	rows, err := res.RowsAffected()
	if err != nil {
		return err
	}
	if rows != 1 {
		return errors.New("User no longer exists")
	}

	return nil
}

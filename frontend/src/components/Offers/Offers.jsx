import { useEffect, useState } from "react";
import axios from 'axios'
import { useNavigate } from "react-router";
import './offers.css'

const Offers = () => {
    const navigate = useNavigate()
    const [offers, setOffers] = useState([])

    //handles on click of logout button
    const handleLogout = () => {
        localStorage.setItem("token", "")
        navigate("/")
    } 
    
    //handles on click of admin portal
    const handleAdmin = () => {
        navigate("/reports")
    }
    
    const handleReport = (reportID, offerID) => {
      const message = document.getElementById("report-offer-button-"+offerID)

      const headers = {
        'Content-Type': 'application/json',
        'Authorization': localStorage.getItem("token")
      }
      const data = {
        reportID: reportID,
        message: message.value,
      }

      //post request to report a user
      axios.post("http://localhost:8080/secure/report-user", data, {headers})
        .then((response) => {
          alert(response.data.response)
        })
        .catch((error) => {
          alert(error.response.data.error)
        })
    }

    const handleCreateOffer = () => {
        const editOfferButton = document.getElementById("create-offer-input")
        const headers = {
            'Content-Type': 'application/json',
            'Authorization': localStorage.getItem("token")
        }
        const data = {
            mealPointsOffer: parseInt(editOfferButton.value)
        }

        //post request to create offer
        axios.post("http://localhost:8080/secure/create-offer", data, {headers})
          .then(() => {
            window.location.reload()
          })
          .catch((error) => {
            alert(error.response.data.error)
          })
    }

    const handleEdit = (offerID, index) => {

        const editOfferButton = document.getElementById("edit-offer-button-"+offerID)
        const headers = {
            'Content-Type': 'application/json',
            'Authorization': localStorage.getItem("token")
        };

        const data = {
            offerID: parseInt(offerID),
            mealPointsOffer: parseInt(editOfferButton.value),
        };

        //put request to update certain offer
        axios.put("http://localhost:8080/secure/update-offer", data, {headers})
          .then(() => {
            const newOffers = [...offers]
            newOffers[index].mealPoints = parseInt(editOfferButton.value) 
            setOffers(newOffers)
          })
          .catch((error) => {
            alert(error.response.data.error)
          })
    }

    //delete request to delete certain offer
    const handleDelete = (offerID, index) => {
        const headers = {
            'Content-Type': 'application/json',
            'Authorization': localStorage.getItem("token"),
        };

        const data = {
            offerID: parseInt(offerID),
        };
        
        axios.delete("http://localhost:8080/secure/delete-offer", {headers, data})
          .then(() => {
            const newOffers = [...offers]
            newOffers.splice(index, 1)
            setOffers(newOffers)
          })
          .catch((error) => {
            alert(error.response.data.error) 
          })
    }

    useEffect(() => {
         const headers = {
            'Content-Type': 'application/json',
            'Authorization': localStorage.getItem("token"),
        };
        
        //get request to get offers, run on init of webpage
        axios.get("http://localhost:8080/secure/get-offers", {headers})
          .then((response) => { 
            setOffers(response.data.offers)
          })
          .catch((error) => {
            //check if unverified email user
            if (error.response.status == 401 && error.response.data.error == "unverified email") {
              alert("Please verify your email first!")
              navigate("/verify-email")
              return
            }
            if (error.response.status == 401){
              alert("Unauthorized")
              navigate("/login")
              return
            }
            alert(error.response.data.error)
          })

    }, [])

    // Create JSX elements for each offer
    const offersJSX = offers?.map((offer, index) => {
        //if users is creator of post, post add delete and edit button
        if (offer.userID == localStorage.getItem("userID")) {
            return (
                <div key={offer.id} className="offer-card">
                    <div className="offer-card-actions">
                        <input
                            id={"report-offer-button-"+offer.id}
                            className="offer-card-report-input"
                            type="text"
                            placeholder="Report Message"
                            
                        />
                        <button className="offer-card-edit-button" onClick={() => handleReport(offer.userID, offer.id)}>Report</button>

                        <input
                            id={"edit-offer-button-"+offer.id}
                            className="offer-card-edit-input"
                            type="number"
                            min="0"
                            max="500"
                            defaultValue={offer.mealPoints}
                        />
                        <button className="offer-card-edit-button" onClick={() => handleEdit(offer.id, index)}>Edit</button>
                        <button className="offer-card-delete-button" onClick={() => handleDelete(offer.id, index)}>Delete</button>
                    </div>
                    <p className="offer-card-meal-points">{offer.mealPoints} Meal Points</p>
                    <p className="offer-card-email">Contact: {offer.email}</p>
                </div>
            )
        }
        else {
            return (
                <div key={offer.id} className="offer-card">
                  <div className="offer-card-actions">
                      <input
                            id={"report-offer-button-"+offer.id}
                            className="offer-card-report-input"
                            type="text"
                            placeholder="Report Message"
                      />
                      <button className="offer-card-edit-button" onClick={() => handleReport(offer.userID, offer.id)}>Report</button>
                  </div>
                    <p className="offer-card-meal-points">{offer.mealPoints} Meal Points</p>
                    <p className="offer-card-email">Contact: {offer.email}</p>
                </div>
            )    
        }
    })
    
    return (
      <div className="offers-container">
        <h1>Washu Meal Points Trader</h1>
        <button className="logout-button" onClick={handleLogout}>Logout</button>
        <button className="logout-button" onClick={handleAdmin}>Admin Portal</button>
        <div className="create-offer-container">
          <input id="create-offer-input" className="create-offer-input" type="number" min="0" max="500" placeholder="0"/>
          <button className="create-offer-button" onClick={handleCreateOffer}>Create offer</button>
        </div>
        <div className="offers-list">
          {offersJSX}
        </div>
      </div>
    
    );
}

export default Offers
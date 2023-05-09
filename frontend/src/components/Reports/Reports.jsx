import { useEffect, useState } from "react";
import axios from 'axios'
import './reports.css'
import { useNavigate } from "react-router";

const Reports = () => {
    const navigate = useNavigate()
    const [reports, setReports] = useState([])

    const handleBanUser = (reportUserID) => {

        
        const headers = {
            'Content-Type': 'application/json',
            'Authorization': localStorage.getItem("token"),
        };
        console.log(reportUserID)
        const data = {
            banUserID: reportUserID,
        }
        
        //put request to ban user
        axios.put(process.env.API_BASE_URL + "/admin/ban-user", data, {headers})
        .then((response) => { 
            alert(response.data.body)
        })
        .catch((error) => {

          if (error.response.status == 401){
            alert("Unauthorized")
            navigate("/login")
            return
          }
          alert(error.response.data.error)
        })
    }
    
    useEffect(() => {
        const headers = {
           'Content-Type': 'application/json',
           'Authorization': localStorage.getItem("token"),
       };
   
       //get request on initial load to get reports
       axios.get(process.env.API_BASE_URL + "/admin/get-reports", {headers})
         .then((response) => { 
           setReports(response.data.reports)
         })
         .catch((error) => {

          if (error.response.status == 401){
            alert("Unauthorized")
            navigate("/offers")
            return
          }
          alert(error.response.data.error)
         })
    }, [])

    //map all reports to element
    const reportsJSX = reports?.map((report, index) => {
      return (
        <div key={index} className="report-card">
            <p className="report-card-email">Reported email: {report.reportEmail}</p>
            <p className="report-card-email">Reporting user email: {report.userEmail}</p>
            <p className="report-card-message">Message: {report.message}</p>
            <button className="ban-button" onClick={() => handleBanUser(report.reportID)}>Ban User</button>
        </div>
      )
    })
    return (
        <div>
            <h1>Reports</h1>
            {reportsJSX}
        </div>
    )
}

export default Reports
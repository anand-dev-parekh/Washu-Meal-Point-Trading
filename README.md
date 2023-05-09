# Washu-Meal-Point-Trading
Full stack web application to trade meal points at Washington University in St Louis (Washu)


---
## Background
At Washington University in St Louis (Washu), many students are above and below on meal points near the end of the semester. This website allows students to connect with other students to trade meal points!


---
## Improvements
1) Regex to only allow wustl emails to register
2) Improve design of frontend
3) Add versitility to backend for admin portal


---
## Test Locally 
*** Make sure to have Node.js, npm, and go installed before continuing
1) clone repo
```console
git clone https://github.com/anand-dev-parekh/Washu-Meal-Point-Trading.git
cd Washu-Meal-Point-Trading
```
2) Set Environment Variables frontend
    a) .env file contains localhost:8000 for base url api

3) Set Environment Variables backend
    a) DBUSER and DBPASS for MySQL user and password
    b) EMAILPASS for gmail app password

4) Install node modules and run frontend
```console
cd frontend
npm install
npm run dev
```
5) Install go modules and run backend
```console
cd backend
go mod download
go run .
```
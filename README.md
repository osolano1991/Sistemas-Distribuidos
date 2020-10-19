POSTMAN

POST
https://osolano-proyecto1.herokuapp.com/shooting/

{
    "id": "400",
    "title": "Title Test",
    "location": "",
    "date": "10/18/2020",
    "incident_area": "UNCA",
    "open_close_location": "Close",
    "target": "random",
    "cause": "Test Cause",
    "summary": "Test summary",
    "fatalities": "0",
    "injured": "0",
    "total_victims": "0",
    "policeman_killed": "0",
    "age": "",
    "employeed_y_n": "",
    "employed_at": "",
    "mental_health_issues": "Unknown",
    "race": "Costa Rican",
    "gender": "Male",
    "latitude": "Male",
    "longitude": "Male"
}

GET ALL
https://osolano-proyecto1.herokuapp.com/shooting/


GET
https://osolano-proyecto1.herokuapp.com/shooting/random
https://osolano-proyecto1.herokuapp.com/shooting/1

PUT
https://osolano-proyecto1.herokuapp.com/shooting/400
https://osolano-proyecto1.herokuapp.com/shooting/random

{
    "id": "400",
    "title": "Title Test - UPDATE",
    "location": "",
    "date": "10/18/2020",
    "incident_area": "UNA UPDATE",
    "open_close_location": "Close UPDATE",
    "target": "random UPDATE",
    "cause": "Cause UPDATE",
    "summary": "Test summary UPDATE",
    "fatalities": "0",
    "injured": "0",
    "total_victims": "0",
    "policeman_killed": "0",
    "age": "",
    "employeed_y_n": "",
    "employed_at": "",
    "mental_health_issues": "Unknown UPDATE",
    "race": "Costa Rican UPDATE",
    "gender": "Male UPDATE",
    "latitude": "Male UPDATE",
    "longitude": "Male UPDATE"
}

DELETE
https://osolano-proyecto1.herokuapp.com/shooting/400

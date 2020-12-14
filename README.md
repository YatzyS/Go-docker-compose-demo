# Amazon webscrapper with docker

- This application demonstrates the use of go-mux for creating a web server, colly for web scrapping and docker and docker-compose to dockerize the application

- This application is made with following tech-stack:
    - Go
    - Docker

- There are 2 folders which has following two services:
    - Scrap - To scrap the data realated to the product from the given link. 
    - Add - To add the data related to the product scrapped from amazon into a JSON file.
- This is a simple application with two services which can help you extract data related to products from amazon which include the name, price, number of reviews, image URL, short description.
- Following is the procedure to run the code:
    - Install go, docker and docker-compose in the system 
    - Clone the project
    - Goto the project directory
    - Just run the following command
        ` docker-compose up`    
    - Open a separate terminal window to test the services.
    - Run the following command in the new terminal:-
        `curl http://localhost:9090/scrap -X POST -d '{"url":"<url goes here>"}'`
    - You will see the product information on the terminal
    - You can also check the data file by opening a terminal into the docker image and serching for a JSON file inside the working directory. 


## BookMyShow

BookMyShow is an online platform for booking tickets for movies and events. It provides the following functionalities:

- Manage user authentication and authorization
- Maintain movie/show information
- Maintain theatre information
- Maintain auditorium and seat information
- Schedule shows for theatres
- Handle online payments and refunds
- Hold reserved seats for 10 minutes

## User/Customer

A user or customer interacts with BookMyShow to perform the following tasks:

- Browse movies and shows
- Search for movies/shows by city, theatre, or movie
- Select a showtime, theatre, and seat type for a movie/show
- Make payment for the selected show
- Receive booking confirmation and ticket details
- Cancel a booking

## Use Cases

Here are the use cases for the interaction between the user/customer and BookMyShow:

1. User/Customer --- (Search movies by city, theatre) ---> BookMyShow
2. User/Customer --- (Select a movie) ---> BookMyShow
3. User/Customer --- (Select a show time) ---> BookMyShow
4. User/Customer --- (Select a seat type and number of seats) ---> BookMyShow
5. BookMyShow --- (Check seat availability) ---> Theatre
6. BookMyShow --- (Calculate the total cost) ---> User/Customer
7. User/Customer --- (Make payment) ---> BookMyShow
8. BookMyShow --- (Handle payment) ---> Payment Gateway
9. Payment Gateway --- (Process payment) ---> User/Customer
10. BookMyShow --- (Book seats) ---> Theatre
11. Theatre --- (Reserve seats for 10 mins) ---> BookMyShow
12. User/Customer --- (Cancel booking) ---> BookMyShow
13. BookMyShow --- (Refund payment) ---> Payment Gateway
14. Payment Gateway --- (Process refund) ---> User/Customer

These use cases cover the major functionalities provided by BookMyShow and the interaction between the user/customer, BookMyShow, and the external payment gateway.

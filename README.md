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


## ER Diagram

```
+-----------------+           +-------------+           +---------------+          +------------+
|     City        |           |    Theatre  |           |  Auditorium   |          |    Show    |
+-----------------+           +-------------+           +---------------+          +------------+
| cityId          |<>-------- | theatreId   |<>-------- | auditoriumId  |<>------- | showId     |
| name            |           | name        |           | name          |          | showTime   |
|                 |           |             |           | features      |<>--------| theatre    |
|                 |           |             |           | seats         |          | screen     |
+-----------------+           +-------------+           | costPerSeat   |          | movie      |
                                                        +---------------+          +------------+
                                                                                          
                                           
+---------------+           +------------+           +-----------------+
|   Movie       |           |   Seat     |           |   Payment       |
+---------------+           +------------+           +-----------------+
| movieId       |<>-------- | seatId     |<>-------- | paymentId       |
| name          |           | seatType   |           | orderId         |
| cost          |           | cost       |           | amount          |
| description   |           +------------+           | paymentStatus   |
| poster        |                                    | paymentMethod   |
| trailer       |                                    +-----------------+
| duration      |
| rating        |
| funFacts      |
| grade         |
+---------------+
```

## Class Diagram

```
+-----------------------+       +-----------------------+       +-----------------------+
|        User           |       |        Admin          |       |        Theatre        |
+-----------------------+       +-----------------------+       +-----------------------+
|                       |       |                       |       | - theatreId: int      |
| + searchMovie()       |       | + addTheatre()        |       | - name: string        |
| + bookSeat()          |       | + addAuditorium()     |       | - city: string        |
| + cancelBooking()     |       | + addShow()           |       | - location: string    |
| + makePayment()       |       | + addMovie()          |       | - auditoriums: list   |
| + requestRefund()     |       | + updateTheatre()     |       | + addAuditorium()     |
+-----------------------+       | + updateAuditorium()  |       | + removeAuditorium()  |
                                | + updateShow()        |       +-----------------------+
                                | + updateMovie()       |
                                | + removeTheatre()     |
                                | + removeAuditorium()  |
                                | + removeShow()        |
                                | + removeMovie()       |
                                +-----------------------+

+-----------------------+       +-----------------------+       +-----------------------+
|      Auditorium       |       |        Seat           |       |         Show          |
+-----------------------+       +-----------------------+       +-----------------------+
| - auditoriumId: int   |       | - seatId: int         |       | - showId: int         |
| - name: string        |       | - type: string        |       | - theatre: Theatre    |
| - features: list      |       | - cost: float         |       | - auditorium: Auditorium |
| - seats: list         |       | - isBooked: bool      |       | - movie: Movie        |
| + addSeat()           |       +-----------------------+       | - timeSlot: datetime  |
| + removeSeat()        |       | + book()              |       | - isAvailable: bool   |
+-----------------------+       | + cancelBooking()     |       | - features: list      |
                                | + reserve()           |       | + bookSeat()          |
                                | + release()           |       | + cancelBooking()     |
                                +-----------------------+       +-----------------------+

+-----------------------+
|        Movie          |
+-----------------------+
| - movieId: int        |
| - name: string        |
| - cost: float         |
| - description: string |
| - poster: image       |
| - trailer: video      |
| - duration: int       |
| - rating: float       |
| - funFacts: string    |
| - grade: string       |
+-----------------------+
| + update()            |
+-----------------------+

```

## List of APIs that project includes -


1. Cities:
- GET /cities: Get a list of all cities available.
- POST /cities: Add a new city to the system.
- GET /cities/{city_id}: Get information about a specific city.
- PUT /cities/{city_id}: Update information about a specific city.
- DELETE /cities/{city_id}: Delete a specific city from the system.

2. Theatres:
- GET /theatres: Get a list of all theatres in a city.
- POST /theatres: Add a new theatre to the system.
- GET /theatres/{theatre_id}: Get information about a specific theatre.
- PUT /theatres/{theatre_id}: Update information about a specific theatre.
- DELETE /theatres/{theatre_id}: Delete a specific theatre from the system.

3. Auditoriums:
- GET /auditoriums: Get a list of all auditoriums in a theatre.
- POST /auditoriums: Add a new auditorium to a theatre.
- GET /auditoriums/{auditorium_id}: Get information about a specific auditorium.
- PUT /auditoriums/{auditorium_id}: Update information about a specific auditorium.
- DELETE /auditoriums/{auditorium_id}: Delete a specific auditorium from a theatre.

4. Seats:
- GET /seats: Get a list of all seats in an auditorium.
- POST /seats: Add a new seat to an auditorium.
- GET /seats/{seat_id}: Get information about a specific seat.
- PUT /seats/{seat_id}: Update information about a specific seat.
- DELETE /seats/{seat_id}: Delete a specific seat from an auditorium.

5. Shows:
- GET /shows: Get a list of all shows in a theatre.
- POST /shows: Add a new show to a theatre.
- GET /shows/{show_id}: Get information about a specific show.
- PUT /shows/{show_id}: Update information about a specific show.
- DELETE /shows/{show_id}: Delete a specific show from a theatre.

6. Movies:
- GET /movies: Get a list of all movies in the system.
- POST /movies: Add a new movie to the system.
- GET /movies/{movie_id}: Get information about a specific movie.
- PUT /movies/{movie_id}: Update information about a specific movie.
- DELETE /movies/{movie_id}: Delete a specific movie from the system.

7. Payments:
- POST /payments: Make a new payment for a booking.
- GET /payments/{payment_id}: Get information about a specific payment.
- PUT /payments/{payment_id}: Update information about a specific payment.
- DELETE /payments/{payment_id}: Cancel a specific payment.

8. Bookings:
- GET /bookings: Get a list of all bookings made by a user.
- POST /bookings: Make a new booking for a show.
- GET /bookings/{booking_id}: Get information about a specific booking.
- PUT /bookings/{booking_id}: Update information about a specific booking.
- DELETE /bookings/{booking_id}: Cancel a specific booking.

9. Features:
- GET /features: Get a list of all features available in an auditorium.
- POST /features: Add a new feature to an auditorium.
- GET /features/{feature_id}: Get information about a specific feature.
- PUT /features/{feature_id}: Update information about a specific feature.
- DELETE /features/{feature_id}: Delete a specific feature from an auditorium.

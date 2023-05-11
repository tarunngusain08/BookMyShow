# BookMyShow

# Use Cases 

+--------------------------------------------------------------------+
|                               User                                 |
+--------------------------------------------------------------------+
| Browse movies and shows                                            |
| Search for movies/shows by city, theatre, or movie                 |
| Select a showtime, theatre, and seat type for a movie/show         |
| Make payment for the selected show                                 |
| Receive booking confirmation and ticket details                    |
| Cancel a booking                                                   |
+--------------------------------------------------------------------+

+--------------------------------------------------------------------+
|                            BookMyShow                              |
+--------------------------------------------------------------------+
| Manage user authentication and authorization                       |
| Maintain movie/show information                                    |
| Maintain theatre information                                       |
| Maintain auditorium and seat information                           |
| Schedule shows for theatres                                        |
| Handle online payments and refunds                                 |
| Hold reserved seats for 10 minutes                                 |
+--------------------------------------------------------------------+


User/Customer --- (Search movies by city, theatre) ---> BookMyShow
User/Customer --- (Select a movie) ---> BookMyShow
User/Customer --- (Select a show time) ---> BookMyShow
User/Customer --- (Select a seat type and number of seats) ---> BookMyShow
BookMyShow --- (Check seat availability) ---> Theatre
BookMyShow --- (Calculate the total cost) ---> User/Customer
User/Customer --- (Make payment) ---> BookMyShow
BookMyShow --- (Handle payment) ---> Payment Gateway
Payment Gateway --- (Process payment) ---> User/Customer
BookMyShow --- (Book seats) ---> Theatre
Theatre --- (Reserve seats for 10 mins) ---> BookMyShow
User/Customer --- (Cancel booking) ---> BookMyShow
BookMyShow --- (Refund payment) ---> Payment Gateway
Payment Gateway --- (Process refund) ---> User/Customer



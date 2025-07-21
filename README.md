# Language- GO

### Go was developed by google in the year 2007   and opensourced in 2009 .

### Why do you think go was developed , even though we have many different languages ?

     Infrastructure where applications were deployed  has changed a lot in the recent times

     Multi core processors became common and using cloud infra with hundred or even thousands of servers  with multiple processors to deploy applications became universal. so infrastructure became much more scalable , dynamic and had more capacity.
     
     However most programming languages didnt help in writing applications that would take advantage  of these infrastructure improvements. in simple words , we have applications that perform one task at a time, but with infrastructure improvements it is now possile to write applications that would execute tasks in parallel.

     For example google drive , Download , Upload , Navigate UI  all happen in parallel or yt watching video , see how many likes , write a comment while watching , see other notified videos in the right tab. but while doing a particular thing do you face any issue with other task ? No right,  so this is a concept of multi threading.

    so each thread is doing one task at a time  and make others running parallel , this makes applications fast but may encounter some issues .
    
    for example a single ticket of a plane is available and three people tries to get it but ony one can buy  it this is conccurency  , so this also has to be handled 

    many languages can  it but the code gets pretty complex and handling and preventing the conccurency issue can be pretty hard , here *GO* comes in to the picture.
   
### Why go ?

    Designed to run multiple cores and build to support concurency 

    concurency in go is cheap and easy 

    Performant Applications 

    Running on scaled and distributed systems 

    Simple and readable syntax like python and effeciency and low level langauge like c++ .
    
    Go is maily used in server side or as  backend language . (Microservices,web applicaitons and database services).
    
    most of the cloud applications like , docker , vault , kubernetes and cockroach db are written in GO.

### Characteristics of GO  : 

    simple syntax , fast build time start up and run , 

    requires  very few resources , 

    compiles easily to single binary code (machine code)

    faster than interpretted languages like python.
    
### Download and working with go :
    
    Navigate to this link to  download the go - https://go.dev/doc/install.
    
    we need main.go ( a standard name where main application is written)
    
    Initialize the project first before you try to run the main file . 

    ```
    go mod init booking-app

    ```
    The above command initiates the go application into a module or project. creates a go.mod file having the version of go.

    Everything in go is a package   , The first statment in go must be  package and the  name of package which our app is part of in our case package main 

    the next command should make go aware that where the entry point is so main function is the entry point for go.
    
    to print a statemnt we cant directly use print command before that we need to install the fmt package and then used the print keyword.  
  

### Lets also work in making a ticket booking system for a cruise (main.go), with the following key functionalities:


**Summary of the App's Workflow:**

1. Introduction & Welcome:
   * The program starts by greeting the user and showing the total number of tickets available (50 in this case) and how many are remaining.


2. User Input:

   * The user is prompted to enter their first name, last name, email address, and the number of tickets they wish to book.

   * After gathering this information, the app validates the input (e.g., ensuring the name is at least 2 characters long, the email is valid, and the number of tickets does not exceed the available tickets).


3. Input Validation:

    * If the input is valid (name, email, and ticket number), the app proceeds with booking.

    * If invalid input is detected, an error message is displayed for the relevant field (name, email, or ticket number).


4. Booking Tickets:

    * Once the input is validated, the bookTickets function is invoked.

    * It reduces the number of remaining tickets by the number of tickets booked.

    * The booking details (first name, last name, email, and number of tickets) are stored in a map and added to the bookings slice.

    * A confirmation message is displayed to the user with the number of tickets remaining.


5. Concurrency: Sending Tickets:

   * The sendTicket function is executed in a goroutine using the go keyword (for concurrency).

   * It simulates sending a confirmation email (using time.Sleep to simulate a delay).

   * Once the goroutine finishes its task (sending the email), it calls wg.Done() to signal the completion of the task.

   * The sync.WaitGroup (wg) ensures the main program waits for all goroutines to finish before exiting (via wg.Wait()).


6. Display of First Names:

   * After booking, the program gathers and prints a list of first names of people who have booked tickets so far.


7. Tickets Sold Out:

    * If the remaining tickets are 0, the program prints a message indicating that the tickets are sold out and encourages the user to return for the next season.

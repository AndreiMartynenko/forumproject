FORUM

Objectives
This project consists in creating a web forum that allows :

- communication between users.
- associating categories to posts.
- liking and disliking posts and comments.
- filtering posts.

SQLite
In order to store the data in your forum (like users, posts, comments, etc.) you will use the database library SQLite.

SQLite is a popular choice as an embedded database software for local/client storage in application software such as web browsers. It enables you to create a database as well as controlling it by using queries.

To structure your database and to achieve better performance we highly advise you to take a look at the entity relationship diagram and build one based on your own database.

You must use at least one SELECT, one CREATE and one INSERT queries.

To know more about SQLite you can check the SQLite page.

Authentication
In this segment the client must be able to register as a new user on the forum, by inputting their credentials. You also have to create a login session to access the forum and be able to add posts and comments.

You should use cookies to allow each user to have only one opened session. Each of this sessions must contain an expiration date. It is up to you to decide how long the cookie stays "alive". The use of UUID is a Bonus task.

Instructions for user registration:

Must ask for email
When the email is already taken return an error response.
Must ask for username
Must ask for password
The password must be encrypted when stored (this is a Bonus task)
The forum must be able to check if the email provided is present in the database and if all credentials are correct. It will check if the password is the same with the one provided and, if the password is not the same, it will return an error response.

Likes and Dislikes
Only registered users will be able to like or dislike posts and comments.

The number of likes and dislikes should be visible by all users (registered or not).

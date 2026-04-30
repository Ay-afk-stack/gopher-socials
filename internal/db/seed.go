package db

import (
	"context"
	"fmt"
	"log"
	"math/rand"
	"strings"

	"github.com/Ay-afk-stack/gopher-socials/internal/store"
	"github.com/jackc/pgx/v5"
)

var usernames = []string{
	"Arjun", "Ramesh", "Sneha", "Anita", "Bikash", "Sushil", "Nisha",
	"Rohit", "Puja", "Manish", "Sarita", "Kiran", "Deepak", "Sunita",
	"Rajesh", "Binita", "Ashok", "Rekha", "Prakash", "Alisha", "Santosh",
	"Sabina", "Gopal", "Kalpana", "Bijay", "Asmita", "Hari", "Rita",
	"Milan", "Sita", "Dipesh", "Karuna", "Ganesh", "Indira", "Roshan",
	"Sharmila", "Krishna", "Laxmi", "Suraj", "Babita",
}

var titles = []string{
	"Getting Started with Go",
	"Understanding REST APIs",
	"Building a CRUD App in Go",
	"Introduction to PostgreSQL",
	"Database Migrations Made Easy",
	"Structuring Your Go Project",
	"Working with JSON in Go",
	"Error Handling Best Practices",
	"Understanding Pointers in Go",
	"Building a Simple Web Server",

	"Introduction to Docker",
	"Containerizing a Go Application",
	"Understanding Microservices",
	"API Design Best Practices",
	"Authentication with JWT",
	"Handling Middleware in Go",
	"Timeouts and Context in Go",
	"Writing Clean Code",
	"Logging in Backend Systems",
	"Environment Variables in Go",

	"Building a Blog API",
	"Pagination in REST APIs",
	"Filtering and Sorting Data",
	"Using UUIDs in Databases",
	"Optimizing Database Queries",
	"Connection Pooling Explained",
	"Handling File Uploads",
	"Testing in Go",
	"Unit vs Integration Testing",
	"Mocking Dependencies in Go",

	"Deploying Go Apps to Production",
	"CI/CD Basics for Developers",
	"Using Nginx as a Reverse Proxy",
	"Scaling Backend Services",
	"Rate Limiting in APIs",
	"Caching Strategies Explained",
	"Securing Your API",
	"Understanding CORS",
	"Monitoring and Metrics",
	"Graceful Shutdown in Go",
}

var contents = []string{
	"This post introduces the basics of Go, covering installation, setup, and writing your first program. It explains core concepts in a simple and practical way for beginners.",
	"Learn how REST APIs work, including request methods, status codes, and best practices for designing scalable and maintainable services.",
	"This guide walks through building a complete CRUD application in Go, from setting up routes to connecting a database and handling requests.",
	"An introduction to PostgreSQL, explaining how relational databases work, basic queries, and how to integrate it with your backend.",
	"Understand how database migrations help manage schema changes safely and consistently across environments.",
	"Learn how to structure a Go project for scalability, including packages, layers, and clean architecture principles.",
	"This article explains how to encode and decode JSON in Go, including tips for working with structs and tags.",
	"Explore error handling patterns in Go and how to write clean, maintainable error logic in your applications.",
	"A beginner-friendly explanation of pointers in Go and how they are used in real-world scenarios.",
	"Build a simple HTTP server in Go using the standard library and understand how routing works.",

	"Get started with Docker and learn how containers simplify development and deployment workflows.",
	"Step-by-step guide to containerizing a Go application using Docker and best practices for building images.",
	"An overview of microservices architecture, including pros, cons, and when to use it.",
	"Learn key principles of API design, including consistency, versioning, and usability.",
	"This post explains how JWT authentication works and how to implement it in your backend.",
	"Understand how middleware works in Go and how to use it for logging, authentication, and more.",
	"Learn how to manage timeouts and cancellations using context in Go applications.",
	"Tips and techniques for writing clean, readable, and maintainable code in any project.",
	"Learn how logging helps debug and monitor backend systems effectively.",
	"Understand how to use environment variables to manage configuration in Go apps.",

	"Build a complete blog API with endpoints for creating, reading, updating, and deleting posts.",
	"Learn how to implement pagination in APIs to efficiently handle large datasets.",
	"This article explains how to filter and sort data in your API queries.",
	"Understand the benefits of using UUIDs instead of integers in databases.",
	"Tips for optimizing database queries to improve performance and reduce load.",
	"Learn how connection pooling works and why it is important for scalable systems.",
	"Step-by-step guide to handling file uploads in a backend service.",
	"Introduction to testing in Go, including writing basic test cases.",
	"Understand the difference between unit testing and integration testing.",
	"Learn how to mock dependencies in Go for better test isolation.",

	"Guide to deploying Go applications to production environments safely.",
	"Introduction to CI/CD and how it improves development workflows.",
	"Learn how to use Nginx as a reverse proxy for your backend services.",
	"Understand strategies for scaling backend systems effectively.",
	"Implement rate limiting to protect your API from abuse.",
	"Explore caching strategies to improve performance and reduce latency.",
	"Learn best practices for securing your API endpoints.",
	"Understand how CORS works and how to configure it properly.",
	"Introduction to monitoring and metrics for backend systems.",
	"Learn how to implement graceful shutdown in Go applications.",
}

var tags = []string{
	"go",
	"golang",
	"backend",
	"api",
	"rest",
	"http",
	"web",
	"server",
	"database",
	"postgres",
	"sql",
	"orm",
	"migrations",
	"docker",
	"containers",
	"devops",
	"microservices",
	"architecture",
	"design",
	"jwt",
	"authentication",
	"authorization",
	"security",
	"middleware",
	"context",
	"timeout",
	"logging",
	"monitoring",
	"testing",
	"unit-testing",
	"integration-testing",
	"mocking",
	"performance",
	"optimization",
	"caching",
	"scaling",
	"nginx",
	"deployment",
	"ci-cd",
	"configuration",
}

var commentContents = []string{
	"Great post, really helped me understand the basics.",
	"This was super clear and easy to follow.",
	"Thanks for sharing this, learned something new today.",
	"Nice explanation, especially the examples.",
	"I was struggling with this topic, this helped a lot.",
	"Well written and very informative.",
	"Can you do a follow-up on this topic?",
	"I like how you broke things down step by step.",
	"This clarified a lot of my doubts.",
	"Simple and to the point, great job.",

	"I tried this and it worked perfectly.",
	"Good overview, but would love more advanced examples.",
	"This is exactly what I was looking for.",
	"Very helpful for beginners like me.",
	"Clean explanation, easy to understand.",
	"Awesome content, keep it coming.",
	"Helped me fix a bug in my project.",
	"I appreciate the practical approach.",
	"Nicely structured and explained.",
	"This saved me a lot of time.",

	"Could you explain this part in more detail?",
	"I think there’s a small typo, but overall great post.",
	"Would love to see more posts like this.",
	"This makes things much clearer.",
	"Great introduction to the topic.",
	"I enjoyed reading this.",
	"Very useful and well explained.",
	"Looking forward to more content from you.",
	"This was a quick and helpful read.",
	"Exactly what I needed, thanks!",
}

func Seed(store store.Storage) {
	ctx := context.Background()

	users := generateUsers(100)

	var tx pgx.Tx

	for _, user := range users {
		if err := store.Users.Create(ctx, tx, user); err != nil {
			log.Printf("error seeding users: %s", err)
			return
		}
	}

	posts := generatePosts(100, users)

	for _, post := range posts {
		if err := store.Posts.Create(ctx, post); err != nil {
			log.Printf("error seeding posts: %s", err)
			return
		}
	}

	comments := generateComments(100, users, posts)

	for _, comment := range comments {
		if err := store.Comments.Create(ctx, comment); err != nil {
			log.Printf("error seeding posts: %s", err)
			return
		}
	}
}

func generateUsers(num int) []*store.User {
	users := make([]*store.User, num)

	for i := 0; i < num; i++ {
		users[i] = &store.User{
			Username: strings.ToLower(usernames[i%len(usernames)] + fmt.Sprintf("%d", i)),
			Email:    strings.ToLower(usernames[i%len(usernames)]+fmt.Sprintf("%d", i)) + "@example.com",
		}
		users[i].Password.Set("1234567890")
	}

	return users
}

func generatePosts(num int, users []*store.User) []*store.Post {
	posts := make([]*store.Post, num)

	for i := 0; i < num; i++ {
		user := users[rand.Intn(len(usernames))]

		posts[i] = &store.Post{
			Title:   titles[rand.Intn(len(titles))],
			Content: contents[rand.Intn(len(contents))],
			UserID:  user.ID,
			Tags: []string{
				tags[rand.Intn(len(tags))],
				tags[rand.Intn(len(tags))],
			},
		}
	}

	return posts
}

func generateComments(num int, users []*store.User, posts []*store.Post) []*store.Comment {
	comments := make([]*store.Comment, num)

	for i := 0; i < num; i++ {
		user := users[rand.Intn(len(users))]
		post := posts[rand.Intn(len(posts))]

		comments[i] = &store.Comment{
			PostID:  post.ID,
			UserID:  user.ID,
			Content: commentContents[rand.Intn(len(commentContents))],
		}
	}

	return comments
}

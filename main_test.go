package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
	"text/template"
)

// TestProfileHandler tests the profile handler for the root route
func TestProfileHandler(t *testing.T) {
	// Initialize the profile data
	profile := Profile{
		Biography: "A highly motivated and adaptable DevOps professional...",
		Skills: []string{
			"CI/CD Tools: Jenkins, Git, GitHub Actions, ArgoCD",
			"Containerization & Orchestration: Docker, Kubernetes",
			"Cloud Platforms: AWS (EC2, VPC, EKS, ECS, CloudWatch, S3, Lambda, Route53, RDS)",
			"Infrastructure as Code (IaC): Terraform, Ansible",
			"Monitoring & Logging: Prometheus, Grafana",
			"Scripting Languages: Bash, Python",
			"Version Control: Git, GitHub",
		},
		Education: "Bachelor of Technology (B.Tech) – Galgotias University...",
		Project: "Self-Driving Car using Raspberry Pi (Great Circle Algorithm + Backtracking)",
	}

	// Create a response recorder to simulate the HTTP request and capture the response
	rr := httptest.NewRecorder()

	// Create a request to the root route
	req, err := http.NewRequest("GET", "/", nil)
	if err != nil {
		t.Fatal(err)
	}

	// Parse the template file
	tmpl, err := template.ParseFiles("profile.html")
	if err != nil {
		t.Fatalf("Error parsing template: %v", err)
	}

	// Handler function to render the profile
	handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		err := tmpl.Execute(w, profile)
		if err != nil {
			http.Error(w, "Unable to load template", http.StatusInternalServerError)
		}
	})

	// Serve the request and capture the response
	handler.ServeHTTP(rr, req)

	// Check if the response status is OK (200)
	if rr.Code != http.StatusOK {
		t.Errorf("Expected status code 200, but got %d", rr.Code)
	}

	// Check if the response body contains some expected content
	if rr.Body.String() == "" {
		t.Errorf("Expected non-empty response body, but got empty")
	}
}
package main

import (
	"html/template"
	"net/http"
)

// Profile struct holds the information to be displayed
type Profile struct {
	Biography string
	Skills    []string
	Education string
	Project   string
}

func main() {
	// Profile data to pass to the HTML template
	profile := Profile{
		Biography: "A highly motivated and adaptable DevOps professional with hands-on experience gained through practical projects, leveraging CI/CD pipelines, cloud infrastructure, and containerization. Skilled in applying DevOps tools and methodologies to optimise workflows, enhance scalability, and drive operational efficiency. Eager to contribute to building automated, scalable deployment pipelines, managing cloud infrastructure, and monitoring systems for peak performance. Ready to thrive in a dynamic team environment while continuously learning and growing within the DevOps domain.",
		Skills: []string{
			"CI/CD Tools: Jenkins, Git, GitHub Actions, ArgoCD",
			"Containerization & Orchestration: Docker, Kubernetes",
			"Cloud Platforms: AWS (EC2, VPC, EKS, ECS, CloudWatch, S3, Lambda, Route53, RDS)",
			"Infrastructure as Code (IaC): Terraform, Ansible",
			"Monitoring & Logging: Prometheus, Grafana",
			"Scripting Languages: Bash, Python",
			"Version Control: Git, GitHub",
		},
		Education: "Bachelor of Technology (B.Tech) – Galgotias University\nAugust 2015 – May 2019\nMajor: Computer Science Engineering with Specialisation in Cloud Computing and Virtualisation\nRelevant Coursework: Cloud Architecture, Containerization, Networking, Virtualisation\nAchievements: Best Interviewee, Runner Up in Product Week",
		Project: "Project: Self-Driving Car using Raspberry Pi (Great Circle Algorithm + Backtracking)",
	}

	// Parse the profile HTML template
	tmpl := template.Must(template.ParseFiles("profile.html"))

	// Serve the profile page
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		err := tmpl.Execute(w, profile)
		if err != nil {
			http.Error(w, "Unable to load template", http.StatusInternalServerError)
		}
	})

	// Start the web server
	http.ListenAndServe(":8080", nil)
}
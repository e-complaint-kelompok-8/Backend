package adminai

import (
	"capstone/entities"
	adminai "capstone/repositories/admin_ai"
	"capstone/repositories/models"
	"context"
	"encoding/json"
	"fmt"
	"os"
	"time"

	"github.com/google/generative-ai-go/genai"
	"google.golang.org/api/option"
)

func NewCustomerService(ai adminai.AISuggestionRepositoryInterface) *AISuggestionService {
	return &AISuggestionService{aiResponseRepo: ai}
}

type AISuggestionServiceInterface interface {
	SaveAISuggestion(adminID, complaintID int, request, response string) error
	GetAISuggestion(request string, complaint entities.Complaint) (string, error)
}

type AISuggestionService struct {
	aiResponseRepo adminai.AISuggestionRepositoryInterface
}

func (service *AISuggestionService) SaveAISuggestion(adminID, complaintID int, request, response string) error {
	aiSuggestion := models.AISuggestion{
		AdminID:     adminID,
		ComplaintID: complaintID,
		Request:     request,
		Response:    response,
		CreatedAt:   time.Now(),
	}
	return service.aiResponseRepo.Create(aiSuggestion)
}

func (service *AISuggestionService) GetAISuggestion(request string, complaint entities.Complaint) (string, error) {
	// Panggil API Generative AI seperti di contoh chatbot
	ctx := context.Background()
	apiKey := os.Getenv("TOKEN_AI")
	client, err := genai.NewClient(ctx, option.WithAPIKey(apiKey))
	if err != nil {
		return "", err
	}
	defer client.Close()

	// Tambahkan detail pengaduan ke dalam konteks permintaan
	genAIParts := []genai.Part{
		genai.Text("Halo, saya adalah Lapi, anggota tim Laporin, sebuah aplikasi untuk membantu pengaduan masyarakat."),
		genai.Text("Detail pengaduan:"),
		genai.Text(fmt.Sprintf("Judul: %s", complaint.Title)),
		genai.Text(fmt.Sprintf("Tanggal: %s", complaint.CreatedAt.Format("02 Jan 2006"))),
		genai.Text(fmt.Sprintf("Alamat: %s", complaint.Location)),
		genai.Text(fmt.Sprintf("Deskripsi: %s", complaint.Description)),
		genai.Text("Permintaan dari Admin: " + request),
		genai.Text("Berikan saran yang spesifik, jelas untuk pengaduan yang diberikan user"),
	}

	model := client.GenerativeModel("gemini-1.5-flash")
	model.SetTemperature(0.7)

	response, err := model.GenerateContent(ctx, genAIParts...)
	if err != nil {
		return "", err
	}
	aiResponse := response.Candidates[0].Content.Parts[0]
	aiResponseString, err := json.Marshal(aiResponse)
	if err != nil {
		return "", err
	}
	return string(aiResponseString), nil
}

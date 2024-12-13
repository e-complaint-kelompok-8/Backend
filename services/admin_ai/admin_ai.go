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
	SaveAISuggestion(adminID, complaintID int, request, response string) (entities.AISuggestion, error)
	GetAISuggestion(request string, complaint entities.Complaint) (string, error)
	GetAISuggestionByID(id string) (entities.AISuggestion, error)
	GetFollowUpAISuggestion(followUpQuery string, aiSuggestion entities.AISuggestion) (string, error)
	GetAllAISuggestions(adminID int) ([]entities.AISuggestion, error)
}

type AISuggestionService struct {
	aiResponseRepo adminai.AISuggestionRepositoryInterface
}

func (service *AISuggestionService) SaveAISuggestion(adminID, complaintID int, request, response string) (entities.AISuggestion, error) {
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
		genai.Text("Kamu adalah Jarvis"),
		genai.Text(DeskripsiChatBot),
		genai.Text("Halo, saya adalah Admin tim Laporin, yang bertugas dalam memberikan tanggapan kepada pengaduan masyarakat."),
		genai.Text("Detail pengaduan user:"),
		genai.Text(fmt.Sprintf("Judul: %s", complaint.Title)),
		genai.Text(fmt.Sprintf("Tanggal: %s", complaint.CreatedAt.Format("02 Jan 2006"))),
		genai.Text(fmt.Sprintf("Alamat: %s", complaint.Location)),
		genai.Text(fmt.Sprintf("Deskripsi: %s", complaint.Description)),
		genai.Text("Permintaan dari Admin: " + request),
		genai.Text("Berikan saran yang spesifik, jelas pada pengaduan yang diberikan oleh user"),
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

func (service *AISuggestionService) GetAISuggestionByID(id string) (entities.AISuggestion, error) {
	return service.aiResponseRepo.GetByID(id)
}

func (service *AISuggestionService) GetFollowUpAISuggestion(followUpQuery string, aiSuggestion entities.AISuggestion) (string, error) {
	ctx := context.Background()
	apiKey := os.Getenv("TOKEN_AI")
	client, err := genai.NewClient(ctx, option.WithAPIKey(apiKey))
	if err != nil {
		return "", err
	}
	defer client.Close()

	// Tambahkan respons AI sebelumnya ke dalam konteks permintaan
	genAIParts := []genai.Part{
		genai.Text("Kamu adalah Jarvis"),
		genai.Text(DeskripsiChatBot),
		genai.Text("Halo, saya adalah Admin tim Laporin, yang bertugas dalam memberikan tanggapan kepada pengaduan masyarakat."),
		genai.Text("Jawaban sebelumnya: " + aiSuggestion.Response),
		genai.Text("Pertanyaan lanjutan dari admin: " + followUpQuery),
		genai.Text("Berikan jawaban yang lebih rinci berdasarkan konteks sebelumnya."),
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

func (service *AISuggestionService) GetAllAISuggestions(adminID int) ([]entities.AISuggestion, error) {
	return service.aiResponseRepo.GetAllByAdminID(adminID)
}

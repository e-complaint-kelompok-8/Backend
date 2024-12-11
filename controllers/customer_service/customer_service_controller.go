package customerservice

import (
	"capstone/controllers/customer_service/request"
	"capstone/controllers/customer_service/response"
	customerservice "capstone/services/customer_service"
	"context"
	"encoding/json"
	"net/http"
	"os"
	"strings"

	"github.com/google/generative-ai-go/genai"
	"github.com/labstack/echo/v4"
	"google.golang.org/api/option"
)

type CustomerServiceController struct {
	customerService customerservice.CustomerServiceInterface
}

func NewCustomerServiceController(cs customerservice.CustomerServiceInterface) *CustomerServiceController {
	return &CustomerServiceController{customerService: cs}
}

func (controller *CustomerServiceController) ChatbotQueryController(c echo.Context) error {
	// Ambil user_id dari JWT di context
	userID, ok := c.Get("user_id").(int)
	if !ok {
		return c.JSON(http.StatusUnauthorized, map[string]interface{}{
			"message": "User not authorized",
		})
	}

	// Ambil data user terkait
	user, err := controller.customerService.GetUserByID(userID)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": "Failed to get user data"})
	}

	request := request.RequestCS{}
	// Ambil pertanyaan dari body request
	if err := c.Bind(&request); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"message": "Invalid request format"})
	}

	// Normalize query
	query := strings.ToLower(strings.TrimSpace(request.Query))

	// Check for static response
	if responseCS, exists := staticResponses[query]; exists {
		return response.SuccessResponse(c, user, request.Query, responseCS)
	}

	// Siapkan konteks dan client AI
	ctx := context.Background()
	apiKey := os.Getenv("TOKEN_AI")
	client, err := genai.NewClient(ctx, option.WithAPIKey(apiKey))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": "Failed to initialize AI client"})
	}
	defer client.Close()

	// Tentukan model AI yang digunakan
	model := client.GenerativeModel("gemini-1.5-flash")
	model.SetTemperature(0.7)

	// Tambahkan konteks aplikasi Laporin
	genAIParts := []genai.Part{
		genai.Text("Halo, saya adalah Lapi, anggota tim Laporin, sebuah aplikasi untuk membantu pengaduan masyarakat."),
		genai.Text(DeskripsiLaporin),
		genai.Text(CaraMengajukanPengaduan),
		genai.Text(CaraMelihatStatusPengaduan),
		genai.Text(CaraMembatalkanPengaduan),
		genai.Text(CaraMembacaBeritaDanPengumuman),
		genai.Text(DiluarTopik),
		genai.Text("Pertanyaan dari user: " + request.Query),
		genai.Text("Berikan jawaban yang spesifik, jelas, dan terkait dengan layanan Laporin."),
	}

	// Kirim permintaan ke model AI
	resp, err := model.GenerateContent(ctx, genAIParts...)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": "Failed to get AI response"})
	}

	// Ambil respons AI
	aiResponse := resp.Candidates[0].Content.Parts[0]

	aiResponseString, err := json.Marshal(aiResponse)
	// cleanedResponse := cleanAIResponse(string(aiResponseString))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Failed to marshal care suggestion"})
	}

	// Simpan respons ke database
	err = controller.customerService.SaveAIResponse(userID, request.Query, string(aiResponseString))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": "Failed to save AI response"})
	}

	// Kirimkan respons ke user
	return response.SuccessResponse(c, user, request.Query, string(aiResponseString))
}

// func cleanAIResponse(response string) string {
// 	// Hapus tanda ** dan \n
// 	response = strings.ReplaceAll(response, "**", "")
// 	response = strings.ReplaceAll(response, "\\n", "\n")
// 	return response
// }

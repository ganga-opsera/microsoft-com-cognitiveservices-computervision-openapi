package main

import (
	"github.com/computer-vision-client/mcp-server/config"
	"github.com/computer-vision-client/mcp-server/models"
	tools_areaofinterest "github.com/computer-vision-client/mcp-server/tools/areaofinterest"
	tools_describe "github.com/computer-vision-client/mcp-server/tools/describe"
	tools_ocr "github.com/computer-vision-client/mcp-server/tools/ocr"
	tools_models "github.com/computer-vision-client/mcp-server/tools/models"
	tools_tag "github.com/computer-vision-client/mcp-server/tools/tag"
	tools_analyze "github.com/computer-vision-client/mcp-server/tools/analyze"
	tools_detect "github.com/computer-vision-client/mcp-server/tools/detect"
)

func GetAll(cfg *config.APIConfig) []models.Tool {
	return []models.Tool{
		tools_areaofinterest.CreateGetareaofinterestTool(cfg),
		tools_describe.CreateDescribeimageTool(cfg),
		tools_ocr.CreateRecognizeprintedtextTool(cfg),
		tools_models.CreateAnalyzeimagebydomainTool(cfg),
		tools_tag.CreateTagimageTool(cfg),
		tools_analyze.CreateAnalyzeimageTool(cfg),
		tools_detect.CreateDetectobjectsTool(cfg),
		tools_models.CreateListmodelsTool(cfg),
	}
}

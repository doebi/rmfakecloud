/*
 * MyScript IInk Batch Mode
 *
 * Service to recognize myscript iink in batch mode
 *
 * API version: 1.1.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package generated

type MathConfiguration struct {
	Solver *SolverConfiguration `json:"solver,omitempty"`
	Margin *MarginConfiguration `json:"margin,omitempty"`
	CustomGrammarId string `json:"customGrammarId,omitempty"`
	CustomGrammarContent string `json:"customGrammarContent,omitempty"`
	UndoRedo *UndoRedoConfiguration `json:"undo-redo,omitempty"`
	SessionTime int32 `json:"session-time,omitempty"`
	Eraser *EraserConfiguration `json:"eraser,omitempty"`
	RecognitionTimeout int32 `json:"recognition-timeout,omitempty"`
}

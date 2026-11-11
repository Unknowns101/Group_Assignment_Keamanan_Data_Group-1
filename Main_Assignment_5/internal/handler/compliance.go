package handler

import (
	"aaa_bekend/internal/db"
	"aaa_bekend/internal/lmodel"
	"encoding/json"
	"net/http"
	"time"
)

type UpdateMyDataReq struct {
    Username string `json:"username"`
}

func GetMyData(w http.ResponseWriter, r *http.Request) {
	 user, err := lmodel.UserFromContext(r.Context())
	    if err != nil {
	        http.Error(w, "Unauthorized", http.StatusUnauthorized)
	        return
	    }
	
	    profile, err := db.GetUserByID(r.Context(), user.ID)
	    if err != nil {
	        http.Error(w, "Failed to fetch user data", http.StatusInternalServerError)
	        return
	    }
	
	    w.Header().Set("Content-Type", "application/json")
	    if err := json.NewEncoder(w).Encode(profile); err != nil {
	        return
	    }
}

func UpdateMyData(w http.ResponseWriter, r *http.Request) {
	user, _ := lmodel.UserFromContext(r.Context())

	// TODO

	//TODO

	json.NewEncoder(w).Encode(map[string]string{
		"message":  "Data updated successfully",
		"username": req.Username,
	})
}

func GetMyAuditLog(w http.ResponseWriter, r *http.Request) {
	user, _ := lmodel.UserFromContext(r.Context())

	// TODO

	// TODO

	json.NewEncoder(w).Encode(map[string]interface{}{
		"user":  user.Username,
		"total": len(logs),
		"logs":  logs,
	})
}

func ExportMyData(w http.ResponseWriter, r *http.Request) {
	user, _ := lmodel.UserFromContext(r.Context())

	// TODO

	// TODO

	export := map[string]interface{}{
		"exported_at": time.Now(),
		"profile":     profile,
		"projects":    projects,
		"audit_logs":  logs,
	}

	w.Header().Set("Content-Disposition", "attachment; filename=my_data_export.json")
	json.NewEncoder(w).Encode(export)
}
func DeleteMyAccount(w http.ResponseWriter, r *http.Request) {
	user, _ := lmodel.UserFromContext(r.Context())

	// TODO

	json.NewEncoder(w).Encode(map[string]string{
		"message": "Your account and all data have been permanently deleted",
	})
}

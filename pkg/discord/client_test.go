package discord

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestGetMe(t *testing.T) {
	expectedUser := User{ID: "123", Username: "TestBot"}
	
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/users/@me" {
			t.Errorf("expected path '/users/@me', got '%s'", r.URL.Path)
		}
		if r.Header.Get("Authorization") != "Bot test-token" {
			t.Errorf("expected Authorization header 'Bot test-token', got '%s'", r.Header.Get("Authorization"))
		}
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(expectedUser)
	}))
	defer ts.Close()

	client := NewClient("test-token")
	client.BaseURL = ts.URL
	
	user, err := client.GetMe()
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}
	
	if user.ID != expectedUser.ID || user.Username != expectedUser.Username {
		t.Errorf("expected user %+v, got %+v", expectedUser, user)
	}
}

func TestCreateMessage(t *testing.T) {
	expectedMessage := Message{ID: "456", Content: "Hello!"}
	
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			t.Errorf("expected POST method, got %s", r.Method)
		}
		if r.URL.Path != "/channels/789/messages" {
			t.Errorf("expected path '/channels/789/messages', got '%s'", r.URL.Path)
		}
		
		var req MessageCreateRequest
		if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
			t.Fatalf("failed to decode request body: %v", err)
		}
		if req.Content != "Hello!" {
			t.Errorf("expected content 'Hello!', got '%s'", req.Content)
		}
		
		w.WriteHeader(http.StatusCreated)
		json.NewEncoder(w).Encode(expectedMessage)
	}))
	defer ts.Close()

	client := NewClient("test-token")
	client.BaseURL = ts.URL
	
	msg, err := client.CreateMessage("789", "Hello!")
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}
	
	if msg.ID != expectedMessage.ID || msg.Content != expectedMessage.Content {
		t.Errorf("expected message %+v, got %+v", expectedMessage, msg)
	}
}

func TestGetUser(t *testing.T) {
	expectedUser := User{ID: "111", Username: "Someone"}
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/users/111" {
			t.Errorf("expected path '/users/111', got '%s'", r.URL.Path)
		}
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(expectedUser)
	}))
	defer ts.Close()

	client := NewClient("token")
	client.BaseURL = ts.URL
	user, _ := client.GetUser("111")
	if user.ID != "111" {
		t.Errorf("expected ID 111, got %s", user.ID)
	}
}

func TestGetGuild(t *testing.T) {
	expectedGuild := Guild{ID: "222", Name: "Noble Server"}
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(expectedGuild)
	}))
	defer ts.Close()

	client := NewClient("token")
	client.BaseURL = ts.URL
	guild, _ := client.GetGuild("222")
	if guild.Name != "Noble Server" {
		t.Errorf("expected Noble Server, got %s", guild.Name)
	}
}

func TestDeleteMessage(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodDelete {
			t.Errorf("expected DELETE method, got %s", r.Method)
		}
		w.WriteHeader(http.StatusNoContent)
	}))
	defer ts.Close()

	client := NewClient("token")
	client.BaseURL = ts.URL
	err := client.DeleteMessage("channel1", "msg1")
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}
}

func TestEditMessage(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPatch {
			t.Errorf("expected PATCH method, got %s", r.Method)
		}
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(Message{ID: "msg1", Content: "Edited"})
	}))
	defer ts.Close()

	client := NewClient("token")
	client.BaseURL = ts.URL
	msg, _ := client.EditMessage("chan1", "msg1", "Edited")
	if msg.Content != "Edited" {
		t.Errorf("expected Edited, got %s", msg.Content)
	}
}

func TestGetMeGuilds(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode([]Guild{{ID: "g1", Name: "Guild 1"}})
	}))
	defer ts.Close()

	client := NewClient("token")
	client.BaseURL = ts.URL
	guilds, _ := client.GetMeGuilds()
	if len(guilds) != 1 || guilds[0].Name != "Guild 1" {
		t.Errorf("unexpected guilds: %v", guilds)
	}
}

func TestGetChannel(t *testing.T) {
	expectedChannel := Channel{ID: "chan1", Name: "Lounge"}
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(expectedChannel)
	}))
	defer ts.Close()

	client := NewClient("token")
	client.BaseURL = ts.URL
	ch, err := client.GetChannel("chan1")
	if err != nil || ch.Name != "Lounge" {
		t.Errorf("unexpected result: %v, %v", ch, err)
	}
}

func TestGetMessage(t *testing.T) {
	expectedMsg := Message{ID: "msg1", Content: "Greetings"}
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(expectedMsg)
	}))
	defer ts.Close()

	client := NewClient("token")
	client.BaseURL = ts.URL
	msg, err := client.GetMessage("chan1", "msg1")
	if err != nil || msg.Content != "Greetings" {
		t.Errorf("unexpected result: %v, %v", msg, err)
	}
}

func TestModifyChannel(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPatch {
			t.Errorf("expected PATCH, got %s", r.Method)
		}
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(Channel{ID: "chan1", Name: "Updated"})
	}))
	defer ts.Close()

	client := NewClient("token")
	client.BaseURL = ts.URL
	ch, _ := client.ModifyChannel("chan1", "Updated")
	if ch.Name != "Updated" {
		t.Errorf("expected Updated, got %s", ch.Name)
	}
}

func TestDeleteChannel(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusNoContent)
	}))
	defer ts.Close()

	client := NewClient("token")
	client.BaseURL = ts.URL
	err := client.DeleteChannel("chan1")
	if err != nil {
		t.Errorf("expected no error, got %v", err)
	}
}

func TestGetGuildChannels(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode([]Channel{{ID: "c1", Name: "C1"}})
	}))
	defer ts.Close()

	client := NewClient("token")
	client.BaseURL = ts.URL
	channels, _ := client.GetGuildChannels("g1")
	if len(channels) != 1 {
		t.Errorf("expected 1 channel, got %d", len(channels))
	}
}

func TestErrorHandling(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode(map[string]string{"message": "Not Found"})
	}))
	defer ts.Close()

	client := NewClient("token")
	client.BaseURL = ts.URL
	_, err := client.GetMe()
	if err == nil {
		t.Error("expected error for 404 status, got nil")
	}
}

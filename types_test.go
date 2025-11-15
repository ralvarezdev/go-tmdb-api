package gotmdbapi

import (
	"context"
	"os"
	"testing"
)

// CreateClient creates a new TMDB API client using the TMDB_TOKEN environment variable
//
// Returns:
//
// - *Client: the TMDB API client
// - error: if there was an error creating the client
func CreateClient() (*Client, error) {
	token := os.Getenv("TMDB_API_KEY")
	return NewClient(token)
}

// TestGetMoviesNowPlayingEndpoint tests the GetMoviesNowPlaying endpoint of the TMDB API client
// 
// Parameters:
// 
// - t: the testing.T instance
func TestGetMoviesNowPlayingEndpoint(t *testing.T) {
	// Create the TMDB API client
	client, err := CreateClient()
	if err != nil {
		t.Fatalf("Failed to create TMDB API client: %v", err)
	}

	// Call the GetMoviesNowPlaying method
	response, statusCode, err := client.GetMoviesNowPlaying(context.Background(), "en-US", 1, "")
	if err != nil {
		t.Fatalf("GetMoviesNowPlaying failed with status code %d: %v", statusCode, err)
	}
	
	// Check if the response is not nil
	if response == nil {
		t.Fatal("GetMoviesNowPlaying returned nil response")
	}

	// Check if the results are not empty
	if len(response.Results) == 0 {
		t.Fatal("GetMoviesNowPlaying returned empty results")
	}

	t.Logf("GetMoviesNowPlaying returned %d results", len(response.Results))
}
	
// TestGetMovieDetailsEndpoint tests the GetMovieDetails endpoint of the TMDB API client
// 
// Parameters:
// 
// - t: the testing.T instance
func TestGetMovieDetailsEndpoint(t *testing.T) {
	// Create the TMDB API client
	client, err := CreateClient()
	if err != nil {
		t.Fatalf("Failed to create TMDB API client: %v", err)
	}

	// Call the GetMovieDetails method for a known movie ID (e.g., 550 for Fight Club)
	response, statusCode, err := client.GetMovieDetails(context.Background(), 550, "en-US")
	if err != nil {
		t.Fatalf("GetMovieDetails failed with status code %d: %v", statusCode, err)
	}
	
	// Check if the response is not nil
	if response == nil {
		t.Fatal("GetMovieDetails returned nil response")
	}

	t.Logf("GetMovieDetails returned movie: %s", response.Title)
}

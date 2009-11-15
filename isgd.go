package main

import ("fmt"; "http"; "os"; "flag"; "io";)

func main() {
    flag.Parse(); // Parse command line args

    if flag.NArg() < 1 {
        fmt.Println("Cowardly refusing to shorten a blank URL");
        os.Exit(-1);
    }

    url, error := shortenURL(flag.Arg(0));
    if error != nil {
        fmt.Println(error);
        os.Exit(-1);
    }

    fmt.Println(url);

}

// Developed using the "API" listed here: http://is.gd/api_info.php
func shortenURL (url string) (shortURL string, err os.Error) {
    u := "http://is.gd/api.php?longurl=" + http.URLEscape(url);

    response, _, err := http.Get(u);

    // Make sure we can connect
    if err != nil {
        return
    }

    // Make sure we get a 200 response code
    if response.StatusCode != 200 {
        return "", os.NewError("Could not shorted your URL. Perhaps it was malformed?");
    }

    b, err := io.ReadAll(response.Body);
    response.Body.Close();

    return string(b), nil
}

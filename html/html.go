package html

var ContentOfHTML = `
<!DOCTYPE html>
<html lang="en">
  <head>
    <meta charset="UTF-8" />
    <meta name="viewport" content="width=device-width, initial-scale=1.0" />
    <title>Spotify Authentication Success</title>
    <style>
      body {
        font-family: "Helvetica Neue", Arial, sans-serif;
        background: #0e0e0e;
        margin: 0;
        height: 100vh;
        display: flex;
        flex-direction: column; /* Stack items vertically */
        justify-content: center;
        align-items: center;
        color: #ffffff;
        text-align: center;
        overflow: hidden;
      }

      .spotify {
        background: -webkit-linear-gradient(145deg, #1db954, #1ed760);
        -webkit-background-clip: text;
        -webkit-text-fill-color: transparent;
        letter-spacing: 1px;
        font-size: 48px; /* Larger font size for the title */
        font-weight: bold;
        text-shadow: 2px 2px 4px rgba(0, 0, 0, 0.2);
        align-self: center; /* Center the h1 within the flex container */
      }

      .container {
        background: linear-gradient(145deg, #1db954, #1ed760);
        border-radius: 20px;
        padding: 45px;
        box-shadow: 0 8px 20px rgba(0, 0, 0, 0.25);
        max-width: 450px;
        width: 80%;
        margin: 20px auto; /* Added some margin to the top and bottom */
        transition: transform 0.3s ease-out, box-shadow 0.3s ease-out;
      }

      .container:hover {
        transform: scale(1.03);
        box-shadow: 0 12px 24px rgba(0, 0, 0, 0.3);
      }
      h1 {
        margin-bottom: 25px;
        font-size: 32px; /* Larger font size for the title */
        letter-spacing: 1px;
        color: #ffffff;
        text-shadow: 2px 2px 4px rgba(0, 0, 0, 0.2);
      }
      p {
        font-size: 18px; /* Larger font size for paragraph */
        line-height: 1;
        margin: 10px;
        color: #ffffff; /* Keeping the paragraph white as requested */
      }
    </style>
  </head>
  <body>
    <div class="spotify">Go Spotify CLI</div>
    <div class="container">
      <h2>Authentication Successful</h2>
      <p>Your Spotify authentication was successful.</p>
      <p>You can now close this window and return to the application.</p>
    </div>
  </body>
</html>

`

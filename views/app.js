// Prediction
function createPagePrediction() {
    var opened = window.open("");
    opened.document.write(`<html>
    <head>
        <title>Prediction</title>
        <link rel="stylesheet" href="/css/style.css"/>
        <script src="app.js"></script>
    </head>
    <body>
        <h2>Name Prediction</h2>
        <h2>What's behind your name?</h2>
        <h2>Please enter your name:</h2>
        <form id="form" onsubmit="submitForm(event)">
            <input type="text" id="name"><br><br>
            <input type="submit">
        </form>
    </body></html>`);
}
function submitForm(e) {
    var greetings = document.getElementById("name").value;
    if (greetings == "") {
        alert("Please enter your name")
        e.preventDefault();
        return
    }

    alert("Hi "+greetings+"! Wait a minute...");
    e.preventDefault();

    fetch('/game/v1/name', {
      method: 'POST',
      headers: {
        'Content-Type': 'application/json'
      },
      body: JSON.stringify({
        "Name": greetings,
      })
    }).then(function(response) {
      return response.json();
    }).then(function(data) {
    // Success code goes here
      alert("Here your result:\n"+JSON.stringify(data))
    }).catch(function(err) {
    // Failure
      alert('Error: '+err)
    });
}

// Dare
function createPageDare() {
    var opened = window.open("");
    opened.document.write(`<html>
    <head>
        <title>Dare or Dare</title>
        <link rel="stylesheet" href="/css/style.css"/>
        <script src="app.js"></script>
    </head>
    <body>
        <h2>Do You Dare?</h2>
        <div class="menu">
        <button onclick="checkDare(event)">Yes, I'm!</button>
        </div>
    </body></html>`);
}
function checkDare(e) {
    alert("Wait a minute...");

    e.preventDefault();

    fetch('/game/v1/dare', {
      method: 'GET',
      headers: {
        'Content-Type': 'application/json'
      }
    }).then(function(response) {
      return response.json();
    }).then(function(data) {
    // Success code goes here
      alert("Here your result:\n"+JSON.stringify(data))
    }).catch(function(err) {
    // Failure
      alert('Error: '+err)
    });
}

// Fact
function createPageFact() {
    var opened = window.open("");
    opened.document.write(`<html>
    <head>
        <title>Today Fact</title>
        <link rel="stylesheet" href="/css/style.css"/>
        <script src="app.js"></script>
    </head>
    <body>
        <h2>Today Fact</h2>
        <div class="menu">
        <button onclick="checkFact(event)">Check it out</button>
        </div>
    </body></html>`);
}
function checkFact(e) {
    alert("Wait a minute...");

    e.preventDefault();

    fetch('/game/v1/fact', {
      method: 'GET',
      headers: {
        'Content-Type': 'application/json'
      }
    }).then(function(response) {
      return response.json();
    }).then(function(data) {
    // Success code goes here
      alert("Here your result:\n"+JSON.stringify(data))
    }).catch(function(err) {
    // Failure
      alert('Error: '+err)
    });
}

function showHide() {
    var x = document.getElementById("menu");
    if (x.style.display === "none") {
        x.style.display = "block";
    } else {
      x.style.display = "none";
    }
  }
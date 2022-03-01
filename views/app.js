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
        <div class="subtitle">
            <h2>Who Are You?</h2>
        </div>
        <h4>Please enter your name:</h4>
        <form id="form" onsubmit="submitForm(event)">
            <input type="text" id="name"><br><br>
            <input type="submit">
        </form>
        <div id="result"></div>
    </body></html>`);
}
function submitForm(e) {
    var name = document.getElementById("name").value;
    if (name == "") {
        alert("Please enter your name")
        e.preventDefault();
        return
    }

    alert("Hi "+name+"! Wait a minute...");
    e.preventDefault();

    fetch('/game/predict', {
      method: 'POST',
      headers: {
        'Content-Type': 'application/json'
      },
      body: JSON.stringify({
        "name": name,
      })
    }).then(function(response) {
      return response.json();
    }).then(function(data) {
    // Success code goes here
      let response = JSON.stringify(data)
      let result = JSON.parse(response)
      let age = result.result.predict_age
      let gender = result.result.predict_gender      
      let nat = ""
      const nats = [];

      // check each country name
      for (let i = 0; i < result.result.predict_nationality.country.length; i++) {
        // if theres empty country name, skip it
        if (result.result.predict_nationality.country[i].country_name == "") {
          continue
        }

        // add to array so we can join it with separator
        nats.push(result.result.predict_nationality.country[i].country_name);
        nat = nats.join(" / ");
      }

      // if data country not found, replace with new value
      if (nat == "") {
        nat = "earth"
      }

      alert("Hi "+name+"! Based on your name, we think you are "+age+" years old.\nWe assume that you are a "+gender+" from "+nat+".")  
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
        <div class="subtitle">
          <h2>Do You Dare?</h2>
        </div>
        <div class="menu">
        <button onclick="checkDare(event)">Yes, I'm!</button>
        </div>
    </body></html>`);
}
function checkDare(e) {
    alert("Wait a minute...");

    e.preventDefault();

    fetch('/game/dare', {
      method: 'GET',
      headers: {
        'Content-Type': 'application/json'
      }
    }).then(function(response) {
      return response.json();
    }).then(function(data) {
    // Success code goes here
      let response = JSON.stringify(data)
      let result = JSON.parse(response)
      let dare = result.result.random_activity
      alert("Let's do it!\n"+dare) 
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
        <div class="subtitle">
          <h2>Did You Know?</h2>
        </div>
        <div class="menu">
        <button onclick="checkFact(event)">Check it out!</button>
        </div>
    </body></html>`);
}
function checkFact(e) {
    alert("Wait a minute...");

    e.preventDefault();

    fetch('/game/fact', {
      method: 'GET',
      headers: {
        'Content-Type': 'application/json'
      }
    }).then(function(response) {
      return response.json();
    }).then(function(data) {
    // Success code goes here
    let response = JSON.stringify(data)
    let result = JSON.parse(response)
    let fact = result.result.random_fact
    alert(fact) 
    }).catch(function(err) {
    // Failure
      alert('Error: '+err)
    });
}

// Fun
function createPageFun() {
  var opened = window.open("");
  opened.document.write(`<html>
  <head>
      <title>Today Fun</title>
      <link rel="stylesheet" href="/css/style.css"/>
      <script src="app.js"></script>
  </head>
  <body>
      <div class="subtitle">
        <h2>Have a Nice Day!</h2>
      </div>
      <div class="menu">
      <button onclick="getImg(event)">A feast for the eyes (keep clicking!)</button><br><br>
      <button onclick="checkJoke(event)">LOL</button>
      </div>
      <div class="space"></div>
  </body></html>`);
}
function getImg(e) {
  e.preventDefault();

  fetch('/game/img', {
    method: 'GET',
    headers: {
      'Content-Type': 'application/json'
    }
  }).then(function(response) {
    return response.json();
  }).then(function(data) {
  // Success code goes here
  let response = JSON.stringify(data)
  let result = JSON.parse(response)
  let url = result.result.url
  showImg(url, 311, 311, "Wish you happiness!")
  }).catch(function(err) {
  // Failure
    alert('Error: '+err)
  });
}
function checkJoke(e) {
  alert("Wait a minute...");

  e.preventDefault();

  fetch('/game/joke', {
    method: 'POST',
    headers: {
      'Content-Type': 'application/json'
    },
    body: JSON.stringify({
      "is_safe": false,
    })
  }).then(function(response) {
    return response.json();
  }).then(function(data) {
  // Success code goes here
  let response = JSON.stringify(data)
  let result = JSON.parse(response)
  let cat = result.result.category
  let joke = result.result.random_joke
  let flag = ""
  const flags = [];

  // check flags safe, if it false, get flags that are true and show it to front
  if (!result.result.flags.safe) {
    let key = Object.keys(result.result.flags);
    for (let i = 0; i < key.length; i++) {
        if (key[i]=="safe") {
          continue
        }
      flags.push(key[i])
      flag = flags.join(", ");
    }
    if (flag == "") {
      flag = "STRONG LANGUAGE"
    }
    alert("WARNING!\nCONTAINS "+flag.toUpperCase())
  }

  alert("["+cat+"]"+"\n"+joke) 
  }).catch(function(err) {
  // Failure
    alert('Error: '+err)
  });
}

function showImg(src, width, height, alt) {
  var img = document.createElement("img");
  img.src = src;
  img.width = width;
  img.height = height;
  img.alt = alt;

  // This next line will just add it to the <body> tag
  document.body.appendChild(img);
}

function showHide() {
    var x = document.getElementById("menu");
    if (x.style.display === "none") {
        x.style.display = "block";
    } else {
      x.style.display = "none";
    }
  }
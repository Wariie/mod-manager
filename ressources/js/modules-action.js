
  var myName = "mod-manager";

  var mySecret = "";

  function command(hash, command) {
    var pl = new Object();
    pl.Hash = hash;
    pl.Name = myName;
    pl.Command = command;
    pl.Type = "Command";
    pl.Secret = mySecret;

    sendCommand(pl);
  }

  function checkInput() {
    var event = window.event || event.which;

    if (event.keyCode == 13) {
        event.preventDefault();
        addLine(document.getElementById("textinput").value);
        document.getElementById("textinput").value = "";
    }

    document.getElementById("textinput").style.height = (document.getElementById("textinput").scrollHeight) + "px";
}

function addLine(line) {
    var textNode = document.createTextNode(line);
    document.getElementById("consoletext").appendChild(textNode);
}

function setSecret(s) {
    mySecret = s;
}


function sendCommand(pl) {
    $.ajax({
        type: "POST",
        url: "/cmd",
        data: JSON.stringify(pl),
        dataType: "html",
        success: function(data) {
            $('#textinput').val($('#textinput').val()+"Command - " + pl.Command + " : " + data +  String.fromCharCode(13, 10));
        },
        error: function() {
            alert('Error occured');
        }
    });
}

function startModule() {

    modToStart = $('#modToStart').val();

    var pl = new Object();
    pl.Name = myName;
    pl.Hash = "hub";
    pl.Command = "Start";
    pl.Type = "Command";
    pl.Secret = mySecret;
    pl.Content = modToStart;

    sendCommand(pl);
}
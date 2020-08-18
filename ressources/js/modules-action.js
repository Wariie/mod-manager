
  var myName = "mod-manager";

  function command(hash, command) {
    var pl = new Object();
    pl.Hash = hash;
    pl.Name = myName;
    pl.Command = command;
    pl.Type = "Command";

    var b = JSON.stringify(pl);
    console.log(b);
    $.ajax({
        type: "POST",
        url: "/cmd",
        data: b,
        dataType: "html",
        success: function(data) {
            $("#textinput").append("Command - " + command + " : " + data);
        },
        error: function() {
            alert('Error occured');
        }
    });
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

var myName = "mod-manager";

var mySecret = "";

var mods;

function command(hash, command) {
    var pl = new Object();
    pl.Hash = hash;
    pl.Name = myName;
    pl.Command = command;
    pl.Type = "Command";
    pl.Secret = mySecret;

    sendCommand(pl);
}

function setParams(s, m) {
    mySecret = s;
    mods = m;
}

function sendCommand(pl) {
    $.ajax({
        type: "POST",
        url: "/cmd",
        data: JSON.stringify(pl),
        dataType: "html",
        success: function(data) {
            toggleModal()
            setModalText(pl.Command, pl.Hash, data);
            //$('#textinput').val($('#textinput').val()+"Command - " + pl.Command + " : " + data +  String.fromCharCode(13, 10));
        },
        error: function() {
            alert('Error occured');
        }
    });
}

function setModalText(cmd, mod, txt) {
    var name = ""
    for (m in mods) {
        if(m.Hash == mod) {
            name == m.Name
        }
    }
    var ansi_up = new AnsiUp;

    var html = ansi_up.ansi_to_html(txt);

    $('.modal-card-title').text(cmd + " > " + name);
    $('.modal-content').html("<p>" + html.replace(/\n/g, "</p><p>") + "</p>");
    $('#modalResult').val("Success");
} 

function toggleModal() {
    $('.modal').toggleClass("is-active");
}

function startModule() {

    modToStart = $('#modToStart option:selected').val();

    var pl = new Object();
    pl.Name = myName;
    pl.Hash = "hub";
    pl.Command = "Start";
    pl.Type = "Command";
    pl.Secret = mySecret;
    pl.Content = modToStart;

    sendCommand(pl);
}


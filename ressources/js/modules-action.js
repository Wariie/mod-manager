
var myName = "mod-manager";

var mySecret = "";

var mods;

function searchHashMod(name) {
    for(m in mods) {
        if(mods[m].NAME == name) {
            return mods[m].PK
        }
    }
}

function command(name, command) {
    var pl = new Object();
    pl.Hash = searchHashMod(name);
    pl.Name = myName;
    pl.Command = command;
    pl.Type = "Command";
    pl.Secret = mySecret;

    sendCommand(pl, name);
}

function setParams(s, m) {
    mySecret = s;
    mods = m;
}

function sendCommand(pl, name) {
    $.ajax({
        type: "POST",
        url: "/cmd",
        data: JSON.stringify(pl),
        dataType: "html",
        success: function(data) {
            toggleModal()
            setModalText(pl.Command, name, data);
        },
        error: function() {
            alert('Error occured');
        }
    });
}

function setModalText(cmd, name, txt) {
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

    sendCommand(pl, modToStart);
}

function sendCommands(name) {
    sp = name.split('.')
    Rname = sp[0];
    for (var i = 1; i < sp.length; i++) {
        Rname += "\\." + sp[i]
    }
    c = $('#commands'+Rname+' option:selected').val();
    command(name, c);
}

function sendCCommands(name) {
    sp = name.split('.')
    Rname = sp[0];
    for (var i = 1; i < sp.length; i++) {
        Rname += "\\." + sp[i]
    }
    c = $('#customCommands'+Rname+' option:selected').val();
    command(name, c);
}


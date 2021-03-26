var mods;

function command(name, command) {
    var pl = new Object();
    pl.Hash = name;
    pl.Name = "mod-manager";
    pl.Command = command;
    pl.Type = "Command";

    sendCommand(pl, name);
}

function setParams(m) {
    mods = m;
}

function sendCommand(pl, name) {
    $.ajax({
        type: "POST",
        url: "/mod-manager/api/",
        data: JSON.stringify(pl),
        dataType: "html",
        success: function(data) {
            toggleModal();
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

    if (cmd == "Start") {
        $('#' + name + "_shutdownorstart").toggleClass("is-success").toggleClass("is-danger");
        $('#' + name + "_shutdownorstart").html("Shutdown");
    } else if (cmd == "Shutdown"){
        $('#' + name + "_shutdownorstart").toggleClass("is-success").toggleClass("is-danger");
        $('#' + name + "_shutdownorstart").html("Start");
    }

    $('.modal-card-title').text(cmd + " > " + name);
    $('.modal-content').html("<p>" + html.replace(/\n/g, "</p><p>") + "</p>");
    $('#modalResult').val("Success");
} 

function toggleModal() {
    $('.modal').toggleClass("is-active");
}

function sendCommands(name) {
    sp = name.split('.');
    Rname = sp[0];
    console.log(Rname)
    for (var i = 1; i < sp.length; i++) {
        Rname += "\\." + sp[i];
    }
    c = $('#commands'+Rname+' option:selected').val();
    command(name, c);
}

function sendCCommands(name) {
    sp = name.split('.');
    Rname = sp[0];
    for (var i = 1; i < sp.length; i++) {
        Rname += "\\." + sp[i];
    }
    c = $('#customCommands'+Rname+' option:selected').val();
    command(name, c);
}


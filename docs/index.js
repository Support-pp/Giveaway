//qId = count of the question!
function submitAnswear(qId) {
    reset()

    var aArray =[];
    //Check Question...
    for (i = 1; i < qId + 1; i++) {
        if (document.getElementById("q" + i).value == "") {
            document.getElementById("qx").innerHTML = "No Answear for Question " + i + "?"
            document.getElementById("qx").classList.remove('none');
            return;
        }
    }

    //Check mail
    if (document.getElementById("email").value == "") {
        document.getElementById("qx").innerHTML = "We need your email. It is a important communication tool!"
        document.getElementById("qx").classList.remove('none');
        return;
    }

    //Check accept
    if (!document.getElementById("checkAGB").checked) {
        document.getElementById("qx").innerHTML = "You are younger than 16? Please accept the agb / private policy box if you over 16 Years!"
        document.getElementById("qx").classList.remove('none');
        return;
    }

    //Check google recaptcha
    if (grecaptcha.getResponse() == "") {
        document.getElementById("qx").innerHTML = "You are a bot? If not accept the google recaptcha!"
        document.getElementById("qx").classList.remove('none');
        return;
    }


    for (i = 1; i < qId + 1; i++) {
        aArray[i-1] = document.getElementById("q" + i).value;
    }

    callAPIRequest(aArray)

}
function reset() {
    document.getElementById("qx").classList.add('none');
}


function callAPIRequest(aArray) {
    var settings = {
        "async": true,
        "crossDomain": true,
        "url": "https://giveaway-spp.herokuapp.com/api/new",
        "method": "POST",
        "headers": {
            "content-type": "application/x-www-form-urlencoded",
        },
        
        "data": {
            "g-recaptcha-response": grecaptcha.getResponse(),
            "email": document.getElementById("email").value,
            "fname": document.getElementById("fname").value,
            "answear": "test", 
        }
        
    }

    $.ajax(settings ).done(function (response) {
        console.log(response);
        rep = JSON.parse(response)
        document.getElementById("code-p").innerHTML = rep.code
        $('#basicExampleModal').modal('show')
    }).fail(function(data){
        document.getElementById("qx").innerHTML = "Your are allready on the giveaway list!"
        document.getElementById("qx").classList.remove('none');
    });
}

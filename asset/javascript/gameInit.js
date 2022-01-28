let id = document.getElementById("id_user");
let text= "";
let setp1 = document.getElementById("step1");
let setp2 = document.getElementById("step2");
let err = document.getElementById("err");
let btn = document.getElementById("btn")

function getVal(){
    text = id.valeur
}

function Step(){
    
    if (text != ""){
        console.log("id valider")
        setp1.style.display= "none";
        setp2.style.display = "flex";
       
    }else{
        console.log("id incompler !")
        setp1.style.display= "flex"
        setp2.style.display = "none"
        if (err.innerHTML == ""){
            err.innerHTML = 'Tu dois mettre un nom de joueur ! sale noob!!';
            console.log("Err message id")
        }
    }
}




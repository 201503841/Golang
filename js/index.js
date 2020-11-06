
/*-----------------------------------------------File---------------------------------------------------*/


function Traducir(){
    var codigo = document.getElementById("codigo1").value;
    alert(codigo);

    fetch('../traducir', {
        method: 'POST',
        body: JSON.stringify({"codigo":codigo}),
        headers:{
          'Content-Type': 'application/json'
        }
    }).then(res => res.json())
    .catch(error => console.error('Error:', error))
    .then(response => view(response));
}

function view(response){
    console.log(response)
    document.getElementById("codigosalida").value = response[0].analisis;
    //document.getElementById("txtin").value = '';
}
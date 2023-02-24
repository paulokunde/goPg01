//parametros: 
  // 1º nome do campo do id do objeto
  // 2º url a para o post
  //3º dados json a serem enviados
  function createOrUpdate(campo_id,url, data) {
    var xhr = new XMLHttpRequest();
    let obj;
    xhr.open("POST",url);
    xhr.send(JSON.stringify(data));
    xhr.onreadystatechange = function (evento) {
      if (xhr.readyState === 4) {
        obj = JSON.parse(xhr.responseText);
        $("#" + campo_id ).val(obj.id);
        $("#msg").text("Operação Executada...OK | id:" + obj.id );
        return obj;
      }else{
        $("#msg").text("Erro ao Executar a operação: Status:" + xhr.status);
      }
    }
  }

  function listaMarcas(ul,url,data){
    var xhr = new XMLHttpRequest();
    let obj;
    xhr.open("POST",url);
    xhr.send(JSON.stringify(data));
    xhr.onreadystatechange = function (evento) {
      if (xhr.readyState === 4) {
        marcas = JSON.parse(xhr.responseText);
        marcas.forEach(function(marca){
            console.log(marca.ID);
        })
  
        
      }else{
        //$("#msg").text("Erro ao Executar a operação: Status:" + xhr.status);
        alert("Erro ao Listar");
      }
    }

  }
  
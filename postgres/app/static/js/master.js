$(".btnDetalhe").click(function(){
  var indice = $(this).parent().parent().index();
  var detalhe = $(this).parent().parent().next();
  var status = $(detalhe).css("display");
  if(status == "none")
      $(detalhe).css("display", "table-row");
  else
      $(detalhe).css("display", "none")
});


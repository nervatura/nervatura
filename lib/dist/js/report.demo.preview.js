/*
This file is part of the Nervatura Framework
http://www.nervatura.com
Copyright © 2011-2017, Csaba Kappel
License: LGPLv3
https://raw.githubusercontent.com/nervatura/nervatura/master/LICENSE
*/

window.PDFJS.disableWorker = true;

var pdfDoc = null,
pageNum = 1,
pageRendering = false,
pageNumPending = null,
scale = 1.3,
canvas = document.getElementById('preview_box'),
ctx = canvas.getContext('2d');

function renderPage(num) {
  pageRendering = true;
  pdfDoc.getPage(num).then(function(page) {
    var viewport = page.getViewport(scale);
    canvas.height = viewport.height;
    canvas.width = viewport.width;

    var renderContext = {
      canvasContext: ctx,
      viewport: viewport};
    var renderTask = page.render(renderContext);
    
    renderTask.promise.then(function () {
      pageRendering = false;
      if (pageNumPending !== null) {
        renderPage(pageNumPending);
        pageNumPending = null;}});});
  document.getElementById('page_num').textContent = pageNum.toString();}

function queueRenderPage(num) {
  if (pageRendering) {
    pageNumPending = num;
  } else {
    renderPage(num);}}
    
window.Preview = { 
  onPrevPage: function() {
    if (pageNum <= 1) {
      return;}
    pageNum--;
    queueRenderPage(pageNum);},
    
  onNextPage:function() {
    if(pdfDoc == null){return;}
    if (pageNum >= pdfDoc.numPages) {
      return;}
    pageNum++;
    queueRenderPage(pageNum);},
    
  showReport: function(rdoc) {
    pageNum = 1;
    window.PDFJS.getDocument(rdoc).then(function(pdfDoc_) {
      pdfDoc = pdfDoc_;
      document.getElementById('page_count').textContent = pdfDoc.numPages;
      renderPage(pageNum);});}
}
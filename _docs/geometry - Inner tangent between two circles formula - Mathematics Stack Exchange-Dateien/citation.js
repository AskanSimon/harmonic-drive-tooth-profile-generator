var citation=function(){var e=function(e){var t=e.closest(".js-post-menu"),n=t.attr("data-post-id"),r=t.find(".js-menu-popup-container");if(!(r.find(".js-cite-popup").length>0)){var o=$('<div class="js-cite-popup popup popup d-block ws6 c-auto">    <div class="js-cite-example mb4"></div>    <form class="grid gs8">        <label class="grid--cell"><input type="radio" class="s-radio js-cite-bibtex" name="reftype" checked="checked"/> BibTeX</label>        <label class="grid--cell"><input type="radio" class="s-radio js-cite-amsrefs" name="reftype" /> amsrefs</label>    </form>    <textarea class="js-cite-text s-textarea w100 d-block mt4" rows="9"></textarea>    <button type="button" class="s-btn js-cite-close mt8 mbn4 p2">Close</button></div>');o.find(".js-cite-close").click(function(){o.fadeOutAndRemove()}),o.find("form").on("submit",function(e){e.preventDefault()});var i=o.find(".js-cite-example"),a=o.find(".js-cite-bibtex"),s=o.find(".js-cite-amsrefs"),c=o.find(".js-cite-text");r.append(o),StackExchange.helpers.addSpinner(i),$.ajax({"type":"GET","url":"/posts/"+n+"/citation","dataType":"json","success":function(e){StackExchange.helpers.removeSpinner(),i.html(e.example),c.val(e.bibtex),a.click(function(){c.val(e.bibtex)}),s.click(function(){c.val(e.amsref)})}})}};return{"show":function(t){e(t)}}}();
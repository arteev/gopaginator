gopaginator
==========


Description
-----------

package to create an array url pages for html templating

Installation
------------

This package can be installed with the go get command:

    go get github.com/Arteev/gopaginator
    
Documentation
-------------

Example:

    PagesArray(4, 1002,"?page=%s")

Result:

    [&{<< ?page=1} &{< ?page=3} &{3 ?page=3} &{4 ?page=4} &{5 ?page=5}
    &{... } &{1000 ?page=1000} &{1001 ?page=1001} &{1002 ?page=1002} &{> ?page=5} &{>> ?page=1002}   

Template:

    <ul class="pagination">
        {{range $value := PagesArray 8 9 "?page=%s&cp="}}            
           <li><a href="{{.Url}}{{$.CountOnPage}}">{{.Name}}</a></li>            
        {{end}}
    </ul>
Result:
    
![Paginator](/sample.pages.png)

FAQ
---


License
-------

MIT: http://mattn.mit-license.org/2012

Author
------

Arteev Aleksey

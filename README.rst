OpenGraph
================

Install
-------------

.. code-block:: bash

   $ go get github.com/shibukawa/opengraph

Usage
-----------

It is designed for using with `hero <https://github.com/shiyanhui/hero>`_ template engine.

.. code-block:: none
   :caption: Sample Template

   <%!
   import "github.com/shibukawa/opengraph"
   %>

   <%: func WriteHTML(og opengraph.OpenGraph, writer io.Writer) %>
   <!DOCTYPE html>
   <html>
       <head <%== opengraph.NameSpace %> >
           <meta charset="utf-8">
           <%== og.Write() %>
       </head>
       <body>
       </body>
   </html>

Key parts are the following:

* ``<head <%== opengraph.NameSpace %> >``

  Add namespaces by using ``opengraph.NameSpace`` constant.
  You should use raw output.

* ``<%== og.Write() %>``

  You should use raw output.

The sample is here:

.. code-block:: go
   :caption: Sample Code

   p := opengraph.NewProfile(
       "https://www.example.com",
       "Yoshiki",
       "Yoshiki", "/image/logo.png",
       "shibu.jp diary", "@shibu_jp", "facebookID")
   article := p.Article(
       "/article", "My Favorite Music",
       "test", "/image/blog.png", time.Now())

   var buffer bytes.Buffer
   WriteHTML(article, &buffer)
   fmt.Println(buffer.String())

Sample Output
-----------------

.. code-block:: html

   <!DOCTYPE html>
   <html lang="en">
   <head prefix="og: http://ogp.me/ns# fb: http://ogp.me/ns/fb# article: http://ogp.me/ns/article#" >
       <meta charset="UTF-8">
       <meta property="og:type" content="artcle" />
       <meta content="test" name="description" />
       <meta content="summary" name="twitter:card" />
       <meta content="@shibu_jp" name="twitter:site" />
       <meta content="My Favorite Music" property="og:title" />
       <meta content="https:/www.example.com/article" property="og:url" />
       <meta content="https:/www.example.com/image/blog.png" property="og:image" />
       <meta content="test" property="og:description" />
       <meta content="shibu.jp diary" property="og:site_name" />
       <meta content="facebookID" property="fb:admins" />
       <script type="application/ld+json">
       {
         "@context": "http://schema.org",
         "@type": "NewsArticle",
         "mainEntityOfPage": {
           "@type": "WebPage",
           "@id": "https:/www.example.com/article"
         },
         "headline": "My Favorite Music",
         "image": [
           "https:/www.example.com/image/blog.png"
          ],
         "datePublished": "2017-10-12T00:20:20+09:00",
         "dateModified": "2017-10-12T00:20:20+09:00",
         "author": {

             "@type": "Person",
             "name": "Yoshiki"

         },
          "publisher": {
           "@type": "Organization",
           "name": "Yoshiki",
           "logo": {
             "@type": "ImageObject",
             "url": "https:/www.example.com/image/logo.png"
           }
         },
         "description": "test"
       }
      </script>
      <title>Title</title>
   </head>
   <body>
   </body>
   </html>

Author
------------

Yoshiki Shibukawa

License
------------

https://shibu.mit-license.org/

MIT

Thanks
-----------

* https://github.com/shiyanhui/hero
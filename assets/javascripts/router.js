define([
  'jquery',
  'underscore',
  'backbone',
  'routers/posts_router'
], function($, _, Backbone, PostsRouter){
 
  var initialize = function() {
    new PostsRouter();
    Backbone.history.start({ pushState: true });
  };

  return {
    intialize: initialize
  };
});

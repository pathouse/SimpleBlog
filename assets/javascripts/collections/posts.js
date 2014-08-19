define([
  'jquery',
  'underscore',
  'backbone',
  'models/post'
], function($, _, Backbone, PostModel){
  
  var PostsCollection = Backbone.Collection.extend({
    model: PostModel,
    url: '/api/posts'
  });

  return PostsCollection;
});

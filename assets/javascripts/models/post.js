define([
  'jquery',
  'underscore',
  'backbone'
], function($, _, Backbone){
  
  var PostModel = Backbone.Model.extend({
    
    paramRoom: 'post',

    urlRoot: '/api/posts',

    defaults: {
      title: 'Post Title',
      body: 'Post Body',
      draft: true
    }
  });

  return PostModel;
});

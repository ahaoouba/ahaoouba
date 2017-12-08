/**
 * @license Copyright (c) 2003-2017, CKSource - Frederico Knabben. All rights reserved.
 * For licensing, see LICENSE.md or http://ckeditor.com/license
 */

CKEDITOR.editorConfig = function( config ) {
	// Define changes to default configuration here. For example:
	// config.language = 'fr';
	// config.uiColor = '#AADC6E';
};
CKEDITOR.editorConfig = function( config )     
{     
    
    config.toolbar = 'MyToolbar';
    
    config.toolbar_MyToolbar =     
    [      
        ['Image','Smiley','Link','Unlink','Anchor']
    ];     

    config.image_previewText=' ';
	config.resize_enabled= true ;
	config.height=100;
	config.filebrowserImageUploadUrl= "/article/uploadimg"; //待会要上传的action或servlet
	config.removePlugins = 'elementspath';
	config.removeDialogTabs = 'image:advanced;image:Link';
	config.filebrowserUploadUrl ="/article/uploadfile";
}; 
/**
 * @license Copyright (c) 2003-2016, CKSource - Frederico Knabben. All rights reserved.
 * For licensing, see LICENSE.md or http://ckeditor.com/license
 */

CKEDITOR.editorConfig = function( config ) {
	// Define changes to default configuration here. For example:
	// config.language = 'fr';
	// config.uiColor = '#AADC6E';

	
	config.image_previewText=' ';
	
	config.filebrowserImageUploadUrl= "/article/uploadimg"; //待会要上传的action或servlet
	config.removePlugins = 'elementspath';
	config.removeDialogTabs = 'image:advanced;image:Link'; 
	config.filebrowserUploadUrl ="/article/uploadfile";
};
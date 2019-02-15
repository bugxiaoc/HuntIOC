
  $(document).ready(function(){
        var eleLeft = $('#left').offset().left;
        var isMouseDown = false;
        var borderLen = 4; //左右边框
		
        $('#left').bind({
			
          mousedown:function(e){
			
            var ele = $(this);
            var rightPos = eleLeft + ele.width() + borderLen;
			
            if(rightPos-5 <= e.pageX && e.pageX <= rightPos){
              isMouseDown = true;
 
              //创建遮罩层，防止mouseup事件被其它元素阻止冒泡，导致mouseup事件无法被body捕获，导致拖动不能停止
              var bodyWidth = $('body').width();
              var bodyHeight = $('body').height();
              $('body').append('<div id="mask" style="opacity:0.1;top:0px;left:0px;background-color:green;position:absolute;z-index:9999;width:'+bodyWidth+'px;height:'+bodyHeight+'px;"></div>');
            }
          }
        });

		$('#right').bind({
			
          mousedown:function(e){
			
            var ele =  $('#left');
            var rightPos = eleLeft + ele.width() + borderLen;
			
            if(rightPos-5 <= e.pageX && e.pageX <= rightPos){
              isMouseDown = true;
 
              //创建遮罩层，防止mouseup事件被其它元素阻止冒泡，导致mouseup事件无法被body捕获，导致拖动不能停止
              var bodyWidth = $('body').width();
              var bodyHeight = $('body').height();
              $('body').append('<div id="mask" style="opacity:0.1;top:0px;left:0px;background-color:green;position:absolute;z-index:9999;width:'+bodyWidth+'px;height:'+bodyHeight+'px;"></div>');
            }
          }
        });
 
        $('body').bind({
          mousemove:function(e){
            var ele = $('#left');
			var eler = $('#right');
			
			var baseele = $('#left').width();
			
			
            var rightPos = eleLeft + ele.width() + borderLen;
			
            if(rightPos-5 <= e.pageX && e.pageX <= rightPos){
              ele.css('cursor','e-resize');
              eler.css('cursor','e-resize');
            }else{
              if(!isMouseDown){
                ele.css('cursor','auto');
                eler.css('cursor','auto');
              }
            }
            if(isMouseDown){
              ele.width((e.pageX-eleLeft-borderLen)+'px'); 
              eler.width(eler.width() + (baseele - (e.pageX-eleLeft-borderLen)) + 'px'); 
            }
          },
          mouseup:function(e){
            isMouseDown = false;
            $('#mask').remove();
          }
        });
 
      });
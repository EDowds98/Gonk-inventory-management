package main

import "html/template"

type content struct {
	HtmlToInsert template.HTML
	Script       template.HTML
}

var IndexContent = content{

	template.HTML(`
	<div class="one-half column offset-by-one" style="margin-top: 25%; overflow: hidden;">
		<h1>Forget Everything You Think You Know About Inventory Management.</h1>
		<p>Welcome to Shelf 2.0! Gonk makes inventory management simple. Please browse our website for more information.</p>
		</div>`),
	template.HTML(``),
}

var OurTechContent = content{
	template.HTML(`<div class="two-thirds column offset-by-one" style="margin-top: 25%; overflow: hidden;">
	<h1>Key Components.</h1>
	<h2>The Soul.</h2>
	<img src="../images/uno.png" alt="Arduino Uno" height="360" width="480"/>
	<p>The iconic Arduino Uno is at the core of our design. The 8MHz RISC architecture and 32Kb 
		of Flash memory allows us to leverage the power of the Arduino library. This allows us to prototype rapidly,
		bringing you better products, faster.
	</p>
	<br>
	<h2>Always Connected.</h2>
	<img src="../images/esp.png" alt="ESP-8266" height="360" width="480"/>
	<p>Your MOUSE robot(s) don't exist in isolation. They transmit shelf data 24/7 365 days a year via an on-board ESP8266.
		Your business needs the latest data possible to stay agile and make great decisions. 
		Why settle for less?
	</p>
  </div>`),
	template.HTML(``),
}

var PortalContent = content{
	template.HTML(`<div class="one-half column" style="margin-top: 25%">
	<h4>Please log in to view live updates from your Gonk™ system.</h4>
	<form method="post" action="/login-success">
	  <label for="uname">Username</label><br>
	  <input type="text" id="uname" name="uname"><br>
	  <label for="pwd">Password</label><br>
	  <input type="password" id="pwd" name="pwd">
	  <input type="submit" value="Log in">
	</form>
  </div>`),
	template.HTML(``),
}

var ContactContent = content{
	template.HTML(`<div class="one-half column offset-by-one" style="margin-top: 25%; overflow: hidden;">
	<h2>Contact Us</h2>
	<p>Grid Building<br>
		Heriot Watt University<br>
		Edinburgh<br>
		EH11 4AP
	</p>
   <br>
   <iframe src="https://www.google.com/maps/embed?pb=!1m18!1m12!1m3!1d2236.204155209582!2d-3.3238597838000583!3d55.91116038059597!2m3!1f0!2f0!3f0!3m2!1i1024!2i768!4f13.1!3m3!1m2!1s0x4887c504db3d66d5%3A0x93f0cbd68f093d76!2sGRID%20at%20Heriot-Watt!5e0!3m2!1sen!2suk!4v1643294350670!5m2!1sen!2suk" width="600" height="450" style="border:0;" allowfullscreen="" loading="lazy" style="border-style: solid; border-color: black;"></iframe>
  </div>`),
	template.HTML(``),
}

var LoginContent = content{
	template.HTML(`<table id="mainTable">
	<tr>
	  <th></th>
	  <th>Item 1</th>
	  <th>Item 2</th>
	  <th>Item 3</th>
	  <th>Item 4</th>
	  <th>Item 5</th>
	  <th>Item 6</th>
	  <th>Item 7</th>
	  <th>Item 8</th>
	</tr>
	<tr>
	  <td>Module 1</td>
	  <td></td>
	  <td></td>
	  <td></td>
	  <td></td>
	  <td></td>
	  <td></td>
	  <td></td>
	  <td></td>
	</tr>
	<tr>
	  <td>Module 2</td>
	  <td></td>
	  <td></td>
	  <td></td>
	  <td></td>
	  <td></td>
	  <td></td>
	  <td></td>
	  <td></td>
	</tr>
	<tr>
	  <td>Module 3</td>
	  <td></td>
	  <td></td>
	  <td></td>
	  <td></td>
	  <td></td>
	  <td></td>
	  <td></td>
	  <td></td>
	</tr>
	<tr>
	  <td>Module 4</td>
	  <td></td>
	  <td></td>
	  <td></td>
	  <td></td>
	  <td></td>
	  <td></td>
	  <td></td>
	  <td></td>
	</tr>
	<tr>
	  <td>Module 5</td>
	  <td></td>
	  <td></td>
	  <td></td>
	  <td></td>
	  <td></td>
	  <td></td>
	  <td></td>
	  <td></td>
	</tr>
	<tr>
	<td>Module 6</td>
	<td></td>
	<td></td>
	<td></td>
	<td></td>
	<td></td>
	<td></td>
	<td></td>
	<td></td>
	</tr>
	<tr>
	  <td>Module 7</td>
	  <td></td>
	  <td></td>
	  <td></td>
	  <td></td>
	  <td></td>
	  <td></td>
	  <td></td>
	  <td></td>
	</tr>
	<tr>
	  <td>Module 8</td>
	  <td></td>
	  <td></td>
	  <td></td>
	  <td></td>
	  <td></td>
	  <td></td>
	  <td></td>
	  <td></td>
	</tr>

  </table> 
`),
	template.HTML(`<script src="js/table-update.js" type="text/javascript"></script>`),
}

var AboutUsContent = content{
	template.HTML(`      <div class="two-thirds column offset-by-one" style="margin-top: 25%; overflow: hidden;">
	<h2>Our Story.</h2>
	<section>
	<p>
		From humble beginnings to a Fortune 500 company, we have become
		synonymous with innovating in the Shelf 2.0 space. Every day our Imagineers show up
		at 2pm and work non-stop until 4pm for our customers. We're not just
		an automated inventory tracking solutions company; we're an 
		automated inventory tracking solutions company that cares.
	</p>
	</section>

	<section>
	<p>
		At Gonk we believe in the turtleneck as a conduit of innovation. Please enjoy
		photos of our Imagineers brainstorming the future of shelving.
	</p>
	</section>
	
	<!-- Slideshow container -->
	<div class="slideshow-container">
	
	  <!-- Full-width images with number and caption text -->
	  <div class="mySlides fade">
		<div class="numbertext">1 / 2</div>
		<img src="../images/team1.jpg" style="width:100%">
	  </div>
	
	  <div class="mySlides fade">
		<div class="numbertext">2 / 3</div>
		<img src="../images/team2.jpg" style="width:100%">
	  </div>
	
	  <!-- Next and previous buttons -->
	  <a class="prev" onclick="plusSlides(-1)">&#10094;</a>
	  <a class="next" onclick="plusSlides(1)">&#10095;</a>
	</div>
	<br>
	
	<!-- The dots/circles -->
	<div style="text-align:center">
	  <span class="dot" onclick="currentSlide(1)"></span>
	  <span class="dot" onclick="currentSlide(2)"></span>
	</div> 
	<p><i>"Counts things, avoids people."</i></p>
  </div>`),
	template.HTML(`<script src="js/slideshow.js" type="text/javascript" defer></script>`),
}

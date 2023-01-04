<template>
  <div class="pt-10 pr-10 pb-28">
    <div class=" text-white h-screen">
      <div class="h-3/4 overflow-auto bg-black rounded-md shadow-lg shadow-lime-600">
        <div class="flex justify-center">
          <div class="pt-4">
            <form v-on:submit.prevent="makeWebsiteThumbnail">
              <p> {{ thumbnailUrl }}</p>
            </form>
            Message-ID: 10083817.1075842507824.JavaMail.evans@thyme<br>
            Date: Wed, 13 Dec 2000 09:33:00 -0800 (PST)<br>
            From: drew.fossum@enron.com<br>
            To: ajenkins@lanlaw.com<br>
            Subject: Re: Fossum news rocks Enron stock<br>
            Cc: dhill@bspmlaw.com, dhooker@caes.net, dfossum@enron.com<br>
            Mime-Version: 1.0<br>
            Content-Type: text/plain; charset=us-ascii<br>
            Content-Transfer-Encoding: 7bit<br>
            Bcc: dhill@bspmlaw.com, dhooker@caes.net, dfossum@enron.com<br>
            X-From: Drew Fossum<br>
            X-To: "Alan Jenkins" ajenkins@lanlaw.com<br>
            X-cc: dhill@bspmlaw.com, Dhooker@caes.net, dfossum@enron.com<br>
            X-bcc:<br>
            X-Folder: \Drew_Fossum_Dec2000_June2001_1\Notes Folders\'sent mail<br>
            X-Origin: FOSSUM-D<br>
            X-FileName: dfossum.nsf<br>

            Dammit Jenkins, don't even joke about stuff like that! If the market ever<br>
            realized that I worked here, the stock would go down $10!!! You think Lay is<br>
            taking the first step toward Sec. Treas. in the new Bush admin??? That's the<br>
            rumor I heard a month ago. Hope all goes well for you in sunny hotlanta.<br>
            Its 5 below zero here today so your whole Georgia good ole' boy lifestyle<br>
            thing sounds pretty attractive. AMF DF<br>
            <br>
            <br>
            <br>
            <br>
            "Alan Jenkins" ajenkins@lanlaw.com on 12/13/2000 04:35:48 PM<br>
            To: dhill@bspmlaw.com, Dhooker@caes.net, dfossum@enron.com<br>
            cc:<br>
            <br>
            Subject: Fossum news rocks Enron stock<br>
            <br>
            <br>
            I trust you guys saw this on the wires today:<br>
            <br>
            Enron's stock was off $3-3/16 or 4.1 percent in afternoon trading at $74,<br>
            following an announcement that Jeff Skilling will become its new chief<br>
            executive officer on Feb. 12. An analyst, who asked not to be named, said<br>
            that it had been widely assumed that Drew Fossum would be named CEO;<br>
            therefore, the company's stock traded lower on the news.<br>
          </div>
        </div>
      </div>
    </div>
  </div>
</template>

<script >
import axios from 'axios';

export default {
  name: 'App',

  data() {
    return {
      websiteUrl: '',
      thumbnailUrl: '',
    }
  },

  methods: {
    makeWebsiteThumbnail() {
      axios.get("http://localhost:3000/api/thumbnail")
        .then((response) => {
          this.thumbnailUrl = response.data;
          console.log(this.thumbnailUrl)
        })
        .catch((error) => {
          window.alert(`The API returned an error: ${error}`);
        })
    }
  }
}
</script>

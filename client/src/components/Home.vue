<template>
<div>
    <div class="mb-5">
        <h1 class="mb-3 txt-center">Welcome to Celestial Body Info Website!</h1>
        <h3>
            This website is aimed at providing astronomical information and facts for users to read up on purely for entertainment purposes or educational reasons.
            All data on this site was created by the website owners by researching online trusted resources. Additional data is sourced and pulled directly from
            the NASA database! We plan on continuing to add additional features along with information to this site in order to expand what you can learn and read
            about here.
        </h3>
    </div>
    <div v-if="apod">
        <h1 class="txt-center">Astronomical Photo of the Day</h1>
        <h2 class="txt-center">{{apod.title}}</h2>
        <h3 class="txt-center">By: {{apod.copy_right}}, {{apod.date}}</h3>
        <img class="col-md-12" v-bind:src="apod.hd_url">
        <p class="col-md-10 apod-explanation">{{apod.explanation}}</p>
    </div>
</div>
</template>

<script>
import HomeService from "../services/HomeService";

export default {
    name: "home",
    data() {
        return {
            apod: null
        };
    },
    methods: {
        retrieveAPOD() { // Fetchs all of our stars for the current page.
            HomeService.getAPOD()
                .then(response => {
                    this.apod = response.data;
                    console.log(this.apod);

                })
                .catch(e => {
                    console.log(e);
                })
        }
    },
    mounted() {
        this.retrieveAPOD();
    }
};
</script>

<style>
.txt-center {
    text-align: center;
}

.apod-explanation {
    font-style: italic;
    margin: 0 auto;
    float: none;
    font-size: 23px;
}
</style>

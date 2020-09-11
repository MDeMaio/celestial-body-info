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
    <div v-if="currApod">
        <h1 class="txt-center" id="apodh1">Astronomical Photo of the Day - {{currApod.date}}</h1>
        <h2 class="txt-center">{{currApod.title}}</h2>
        <h3 v-if="currApod.copy_right" class="txt-center">By: {{currApod.copy_right}}</h3>
        <img class="col-md-12" v-bind:src="currApod.hd_url">
        <nav aria-label="Pagination">
            <ul class="pagination mt-2 justify-content-center">
                <li :key="'prev'" class="page-item previous-item" :class="{'disabled':currApodIndex === (apodList.length - 1)}">
                    <button class="page-link" @click="currApodIndex++">Prev</button>
                </li>
                <li :key="'next'" class="page-item next-item" :class="{'disabled':currApodIndex === 0}">
                    <button class="page-link" @click="currApodIndex--">Next</button>
                </li>
            </ul>
        </nav>
        <p class="col-md-10 apod-explanation">{{currApod.explanation}}</p>
    </div>
</div>
</template>

<script>
import HomeService from "../services/HomeService";

export default {
    name: "home",
    data() {
        return {
            apodList: [],
            currApod: null,
            currApodIndex: 0
        };
    },
    methods: {
        retrieveAPOD() { // Fetchs all of our apods for 7 days.
            HomeService.getAPOD()
                .then(response => {
                    this.apodList = response.data;
                    console.log(this.apodList);
                    this.currApod = this.apodList[0];

                })
                .catch(e => {
                    console.log(e);
                })
        }
    },

    watch: {
        "currApodIndex": function (index) {
            if(index < 0 || index >= 7){    // why this would ever happen, idk.
                return;
            }
            this.currApod = this.apodList[index];
            setTimeout(function () {
                document.getElementById("apodh1").scrollIntoView({  // Reposition view to title of apod picture.
                    behavior: 'smooth'
                });
            }, 200);
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

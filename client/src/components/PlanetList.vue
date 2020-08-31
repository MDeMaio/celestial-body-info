<template>
<div>
    <div class="row justify-content-center">
        <div class="col-md-6 list">
            <div class="input-group mb-3">
                <input type="text" class="form-control" placeholder="Search by name" v-model="name" />
                <div class="input-group-append">
                    <button class="btn btn-outline-secondary" type="button" @click="searchName">
                        Search
                    </button>
                </div>
            </div>
        </div>
    </div>
    <div class="row justify-content-center">
        <div class="col-md-6">
            <h4>Planets List</h4>
            <ul class="list-group">
                <li class="list-group-item lgi-pointer" :class="{ active: index == currentIndex }" v-for="(planet, index) in planets" :key="index" @click="setActivePlanet(planet, index)">
                    {{ planet.name }}
                </li>
            </ul>
        </div>
    </div>
    <h1 style="text-align: center; margin-top: 2%; font-size: 50px;" v-if="currentPlanet">{{ currentPlanet.name }}</h1>
    <div class="row mt-3">
        <div class="col-md-4">
            <div v-if="currentPlanet" v-html="generateAttributeHTML(currentPlanet.atmosphere_info,`Atmosphere Info:`, ``, ``)">
            </div>
        </div>
        <div class="col-md-4">
            <div v-if="currentPlanet" v-html="generateAttributeHTML(currentPlanet.physical_info,`Physical Info:`, ``, ``)">
            </div>
        </div>
        <div class="col-md-4">
            <div v-if="currentPlanet" v-html="generateAttributeHTML(currentPlanet.orbital_info,`Orbital Info:`, ``, ``)">
            </div>
        </div>
    </div>
</div>
</template>

<script>
import PlanetService from "../services/PlanetService";

export default {
    name: "planets-list",
    data() {
        return {
            planets: [],
            currentPlanet: null,
            currentIndex: -1,
            name: ""
        };
    },
    methods: {
        retrievePlanets() {
            PlanetService.getAll()
                .then(response => {
                    this.planets = response.data;
                    console.log(response.data);
                })
                .catch(e => {
                    console.log(e);
                });
        },

        refreshList() {
            this.retrievePlanets();
            this.currentPlanet = null;
            this.currentIndex = -1;
        },

        setActivePlanet(planet, index) {
            this.currentPlanet = planet;
            this.currentIndex = index;
        },

        searchName() {
            PlanetService.findByTitle(this.name)
                .then(response => {
                    this.planets = response.data;
                    console.log(response.data);
                })
                .catch(e => {
                    console.log(e);
                });
        },

        capitalizeAttributes(str){  // Capitlize words split by _
            const res = str.split("_");
            let finalVal = "";
            for (let i = 0; i < res.length; i++){
                res[i].charAt(0).toUpperCase();
                finalVal += res[i].charAt(0).toUpperCase() + res[i].slice(1) + " ";
            }

            return finalVal.trim();
        },

            generateAttributeHTML: function(data, category, html, nestedHtml){
            if(this.currentPlanet){
                if(html === ""){
                    html += `<label><strong>` + category + `</strong></label>`
                }
                html += `<ul class="list-group">`;
                for (const prop in data){
                    if(typeof data[prop] === "object" && !Array.isArray(data[prop])){ // Recursive call for nested props.
                         html += `<li class="list-group-item">
                             <strong>` + this.capitalizeAttributes(prop) + `: </strong>`;
                        return this.generateAttributeHTML(data[prop], category, html, nestedHtml) + "</li>";
                    } else if(Array.isArray(data[prop])){   // Handle array of object data.
                        html += `<li class="list-group-item"><strong>` + this.capitalizeAttributes(prop) +": </strong> <ul>";
                        for(const item in data[prop]){
                            html += `<li class="list-group-item">`;
                            for(const nestedItem in data[prop][item]){
                                html += `<strong>` + this.capitalizeAttributes(nestedItem) + `: </strong>` + data[prop][item][nestedItem] + `
                               </br>`;
                            }
                            html += " </li>";
                        }
                        html += "</ul></li>";
                    } 
                    else {  // Default case.
                        nestedHtml = data[prop];
                         html += `<li class="list-group-item">
                             <strong>` + this.capitalizeAttributes(prop) + `: </strong>` + nestedHtml + `
                         </li>`;
                    }
                }
            } else{
                html = "</ul><div></div>"
            }

            return html;
        }
    },
    mounted() {
        this.retrievePlanets();
    }
};
</script>

<style>
.lgi-pointer{
  cursor: pointer;
}
.lgi-pointer:hover{
  background-color: #53a5fc;
}
</style>
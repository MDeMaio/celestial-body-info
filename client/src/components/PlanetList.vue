<template>
<div>
    <div class="row justify-content-center">
        <div class="col-md-6 list">
            <form v-on:submit.prevent @submit="searchName">
                <div class="input-group mb-3">
                    <input ref="name" type="text" class="form-control fs-20" placeholder="Search by name" v-model="name" />
                    <div class="input-group-append">
                        <button class="btn btn-outline-primary ml-1 fs-20" type="submit">
                            Search
                        </button>
                        <button class="btn btn-outline-secondary fs-20" type="button" @click="resetPage">
                            Reset
                        </button>
                    </div>
                </div>
            </form>
        </div>
    </div>
    <div class="row justify-content-center">
        <div class="col-md-6">
            <h4 style="font-size: 35px;">Available Planets</h4>
            <div class="btn-group dropleft mb-2">
                <button type="button" class="btn btn-secondary dropdown-toggle" data-toggle="dropdown" aria-haspopup="true" aria-expanded="false">
                    Planet Type
                </button>
                <div class="dropdown-menu">
                    <button class="dropdown-item" @click="type = 'All'" :class="{'active':(type === 'All')}">All</button>
                    <button class="dropdown-item" @click="type = 'Inner Planet'" :class="{'active':(type === 'Inner Planet')}">Inner Planet</button>
                    <button class="dropdown-item" @click="type = 'Outer Planet'" :class="{'active':(type === 'Outer Planet')}">Outer Planet</button>
                    <button class="dropdown-item" @click="type = 'Exoplanet'" :class="{'active':(type === 'Exoplanet')}">Exoplanet</button>
                </div>
            </div>
            <ul class="list-group">
                <li class="list-group-item lgi-pointer py-1" :class="{ active: index == currentIndex }" v-for="(planet, index) in planets" :key="planet.planet_id" @click="setActivePlanet(planet, index)">
                    {{ planet.name }}
                </li>
            </ul>
            <nav aria-label="Pagination">
                <ul class="pagination mt-2 justify-content-center">
                    <li v-if="loaded" :key="'prev'" :class="{'disabled':currentPage === 1}" class="page-item previous-item">
                        <button class="page-link" @click="currentPage--">Prev</button>
                    </li>
                    <li v-for="page in pageArray" :key="page" class="page-item" :class="{'active':(currentPage === page)}">
                        <button class="page-link" @click="currentPage = page">{{page}}</button>
                    </li>
                    <li v-if="loaded" :key="'next'" :class="{'disabled':currentPage === totalPages}" class="page-item next-item">
                        <button class="page-link" @click="currentPage++">Next</button>
                    </li>
                </ul>
            </nav>
        </div>
    </div>
    <transition name="slide-fade">
        <h1 style="text-align: center; margin-top: 2%; font-size: 50px;" v-if="currentPlanet">{{ currentPlanet.name }}</h1>
    </transition>
    <div class="row">
        <transition name="slide-fade">
            <div class="col-md-4" v-if="currentPlanet">
                <h1>Basic Information</h1>
                <ul class="list-group">
                    <li :key="'alternate_name'" class="list-group-item lgi-colored-space">
                        <p style="font-size: 22px;"><strong>Alternate Name(s):</strong> {{currentPlanet.basic_information.alternate_name}}</p>
                    </li>
                    <li :key="'type'" class="list-group-item lgi-colored-space">
                        <p style="font-size: 22px;"><strong>Planetary Type:</strong> {{currentPlanet.basic_information.type}}</p>
                    </li>
                    <li :key="'number_of_satelites'" class="list-group-item lgi-colored-space">
                        <p style="font-size: 22px;"><strong>Number of Satelites:</strong> {{currentPlanet.basic_information.number_of_satelites > 0 ? currentPlanet.basic_information.number_of_satelites : 0}}</p>
                    </li>
                    <li :key="'star_system'" class="list-group-item lgi-colored-space">
                        <p style="font-size: 22px;"><strong>Star System:</strong> {{currentPlanet.basic_information.star_system}}</p>
                    </li>
                    <li :key="'most_abundant_resource'" class="list-group-item lgi-colored-space">
                        <p style="font-size: 22px;"><strong>Most Abundant Resource:</strong> {{currentPlanet.basic_information.most_abundant_resource}}</p>
                    </li>
                    <li :key="'surface_gravity'" class="list-group-item lgi-colored-space">
                        <p style="font-size: 22px;"><strong>Surface Gravity:</strong> {{currentPlanet.basic_information.surface_gravity}}m/s<sup>2</sup></p>
                    </li>
                </ul>
            </div>
        </transition>
        <transition name="slide-fade">
            <img class="img-center col-md-4 justify-content-right" v-if="currentPlanet" v-bind:src="'data:image/jpeg;base64,' + currentPlanet.image">
        </transition>
        <div class="col-md-4">
        </div>
    </div>
    <div class="row mt-3 justify-content-center">
        <div class="col-md-10">
            <div v-if="currentPlanet">
                <ul class="list-group list-group-flush">
                    <li class="list-group-item mt-2" v-for="(val, index) in currentPlanet.facts" :key="index">
                        <h1 style="font-size: 40px; text-align: center;">{{val.title}}</h1>
                        <p style="font-size: 22px;">{{val.fact}}</p>
                    </li>
                </ul>
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
            name: typeof this.$route.query.name == "undefined" || this.$route.query.name == "All" ? "" : this.$route.query.name,
            pageArray: [],
            currentPage: typeof this.$route.query.page == "undefined" ? 1 : parseInt(this.$route.query.page),
            totalPages: 0,
            type: typeof this.$route.query.type == "undefined" ? "All" : this.$route.query.type,
            loaded: false,
            resetting: false,
            searching: false
        };
    },
    methods: {
        retrievePlanets(page, type, name) { // Fetchs all of our planets for the current page.
            PlanetService.getAll(page, type, name == "" ? "All" : name)
                .then(response => {
                    this.planets = response.data.planets;
                    this.currentPage = page;
                    this.type = type;
                    this.pageArray = this.generatePaginationPageArray(response.data.number_of_documents, 5);
                    this.clearPlanetView();
                })
                .catch(e => {
                    console.log(e);
                }).finally(() => {
                    this.resetting = false; // why here????
                    this.searching = false; // ???????????
                });
        },

        clearPlanetView() { // Refreshes the page to the default state.
            this.currentIndex = -1;
            this.currentPlanet = null;
        },

        setActivePlanet(planet, index) { // Updates currently viewed planet.
            this.currentPlanet = planet;
            this.currentIndex = index;
        },

        searchName() { // This will search for a planet by name and return it, caps specific as of now.
            if (this.name === "") {
                alert("Please enter a planet to search for.");
                this.$refs.name.focus();
                return;
            }

            this.searching = true;
            this.retrievePlanets(1, "All", this.name);
            this.$router.push({ query: Object.assign({}, this.$route.query, { page: 1, type: "All", name: this.name}) });

        },
        validatePageCount(count, total) {
            if (count > total) {
                count = total;
            }
            if (count % 2 == 0) { //ensure odd number
                count--;
            }
            return count;
        },
        generatePaginationPageArray(numOfDocuments, recordsPerPage) { // Fetch our pages in an array so we can iterate over it later to create the pagination list items.
            if (numOfDocuments == 0) { // No paging if no records.
                return;
            }

            const pageArray = [];
            this.totalPages = Math.ceil(numOfDocuments / recordsPerPage); // sets total page count for data
            let pageDisplayCount = this.validatePageCount(5, this.totalPages);
            const mid = Math.ceil(pageDisplayCount / 2); // sets mid of displayed pages
            let startIndex = (this.currentPage + 1) - mid; // set first page to be displayed

            if (this.totalPages <= recordsPerPage) {
                startIndex = 1;
                pageDisplayCount = this.totalPages;
            } else if (this.currentPage < mid) { // account for front half
                startIndex = 1;
            } else if ((this.currentPage + (mid - 1)) > this.totalPages) { //account for back half
                startIndex = this.totalPages - pageDisplayCount + 1;
            }
            for (let i = startIndex; i < startIndex + pageDisplayCount; i++) {
                pageArray.push(i);
            }
            return pageArray;
        },

        resetPage() {
            this.name = '';
            this.resetting = true;
            this.retrievePlanets(1, 'All', 'All');
            this.$router.push(this.$route.path)
        }
    },
    watch: { // Watch for data change in which page the user is currently on, call API to get new data when it changes.
        "currentPage": function () {
            if (this.resetting || this.searching) {
                return;
            }

            this.$router.push({ query: Object.assign({}, this.$route.query, { page: this.currentPage, type: this.type, name: this.name == "" ? "All" : this.name }) });
            this.retrievePlanets(this.currentPage, this.type, this.name);
        },

        "type": function () {
            if (this.resetting || this.searching) {
                return;
            }

            this.$router.push({ query: Object.assign({}, this.$route.query, { page: this.currentPage, type: this.type, name: this.name == "" ? "All" : this.name }) });
            this.name = "";
            if (this.currentPage === 1) { // Otherwise changing current page will take care of the refresh for us.
                this.retrievePlanets(this.currentPage, this.type, this.name);
                return;
            }
            this.currentPage = 1;
        },

        "planets": function (val) {
            if(val.length === 1){
                this.currentPlanet = val[0];
                this.currentIndex = 0;
            }
        }
    },
    mounted() {
        this.retrievePlanets(this.currentPage, this.type, this.name);
        this.loaded = true;
    }
};
</script>

<style>
.lgi-pointer {
    cursor: pointer;
}

.lgi-pointer:hover {
    background-color: #53a5fc;
}

.lgi-colored-space {
    background-color: #fcfbfe;
    color: #1d1135;
}

.img-center {
    display: block;
    margin-left: auto;
    margin-right: auto;
    width: 50%;
}

.fs-20 {
    font-size: 20px;
}

/* Enter and leave animations can use different */
/* durations and timing functions.              */
.slide-fade-enter-active {
    transition: all 1s ease;
}

.slide-fade-leave-active {
    transition: all 1s ease;
}

.slide-fade-enter,
.slide-fade-leave-to

/* .slide-fade-leave-active below version 2.1.8 */
    {
    transform: translatey(300px);
    opacity: 0;
}
</style>

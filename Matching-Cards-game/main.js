var tenantArray=[
    {    'name':'ELON', 'img':'/Users/udaymadhav/Documents/Matching-Cards-Game/images/ElonMusk.jpg?raw=true', },
    {    'name':'RONALDO', 'img':'/Users/udaymadhav/Documents/Matching-Cards-Game/images/CRonaldo.jpg?raw=true',},
    {    'name': 'SACHIN',     'img': '/Users/udaymadhav/Documents/Matching-Cards-Game/images/Sachin.jpg?raw=true',  },
    {    'name': 'MESSI',      'img': '/Users/udaymadhav/Documents/Matching-Cards-Game/images/Messi.png?raw=true',  },
    {    'name': 'MJ',         'img': '/Users/udaymadhav/Documents/Matching-Cards-Game/images/MJ.jpg?raw=true',  },
    {    'name': 'ROCK',       'img': '/Users/udaymadhav/Documents/Matching-Cards-Game/images/Rock.jpg?raw=true',  },
    {    'name': 'KEANU',      'img': '/Users/udaymadhav/Documents/Matching-Cards-Game/images/Keanu.jpg?raw=true',  },
    {    'name': 'PRABHAS',    'img': '/Users/udaymadhav/Documents/Matching-Cards-Game/images/Prabhas.jpg?raw=true',  },
    {    'name': 'IRONMAN',    'img': '/Users/udaymadhav/Documents/Matching-Cards-Game/images/Ironman.jpg?raw=true',  },
    {    'name': 'SUPERMAN',   'img': '/Users/udaymadhav/Documents/Matching-Cards-Game/images/Superman.jpg?raw=true',  },
    {    'name': 'FERRARI',    'img': '/Users/udaymadhav/Documents/Matching-Cards-Game/images/Ferrari.jpg?raw=true',  },
    {    'name': 'SERENA',     'img': '/Users/udaymadhav/Documents/Matching-Cards-Game/images/Serena.jpg?raw=true',  },
];
//Part3 duplicating Array
var tenantGameGrid=tenantArray.concat(tenantArray);
tenantGameGrid.sort(function(){
    return 0.5- Math.random();
})
var tenantDiv=document.querySelector('game-board');
var tenantGame=document.getElementById('game-board');
var tenantSection=document.createElement('section');
tenantSection.setAttribute('class', 'tenantSection');
tenantGame.appendChild(tenantSection);

//part2
for(i= 0; i< tenantGameGrid.length; i++){
    // part2  tenantDiv= card
    var tenantDiv=document.createElement('div');
    //Adding 'css-class' to div-tenantDiv
    tenantDiv.classList.add('card');
    tenantDiv.dataset.name=tenantGameGrid[i].name;
    tenantDiv.style.backgroundImage=`url(${tenantGameGrid[i].img})`;
    tenantSection.appendChild(tenantDiv);
};
var count=0;
//part4 create eventListener when user clicks on any image it should get a border
tenantSection.addEventListener('click', function(event){
    var clicked=event.target;
    if(clicked.nodeName==='SECTION'){
        return " ";
    }
    //part5 user should select only 2 cards at time. Put count=0
    if(count<2){
        count++;
        clicked.classList.add('tenantSelected');
    }
});
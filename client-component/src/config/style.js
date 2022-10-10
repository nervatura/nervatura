import { css } from 'lit';

export default css`
.s1 { float:left; width: 8.33333%; } .s2 { float:left; width: 16.66666%; } .s3 { float:left; width: 24.99999%; }
.s4 { float:left; width: 33.33333%; } .s5 { float:left; width: 41.66666%; } .s6 { float:left; width: 49.99999%; }
.s7 { float:left; width: 58.33333%; } .s8 { float:left; width: 66.66666%; } .s9 { float:left; width: 74.99999%; }
.s10 { float:left; width: 83.33333%; } .s11 { float:left; width: 91.66666%; } .s12 { float:left; width:99.99999%; }
.half,.third,.quarter{ float:left; width:100% }

.tiny { font-size:10px!important; }
.small { font-size:12px!important; }
.medium { font-size:15px!important; }
.large { font-size:18px!important; }

.container { padding: 0.01em 16px; } 
.container::after, .container::before { content: ""; display: table; clear: both; }
.container-small { padding: 0px 8px; }
.padding-tiny { padding: 2px 4px; }
.padding-small { padding: 4px 8px; }
.padding-normal { padding: 8px 16px; }
.padding-large { padding: 12px 24px; }
.padding-xlarge { padding: 16px 32px; }

.left { float: left; }
.right { float: right; }
.align-left { text-align: left; }
.align-right { text-align: right; }
.justify { text-align: justify; }
.center { text-align: center; }
.centered { margin: auto; }
.margin0 { margin: 0!important; }
.top {
  vertical-align: top!important;
}

.bold { font-weight: bold; }
.italic { font-style: italic; }

.section { padding-top: 16px!important; padding-bottom: 16px!important; }
.section-small { padding-top: 8px!important; padding-bottom: 8px!important; }
.section-small-top { padding-top: 8px!important; }
.section-small-bottom { padding-bottom: 8px!important; }
.block { display: block; width: 100%; }
.full { display: block; width: 100%; }

.hide{ display: none!important; }
.show{ display: block!important; }

.row {
  display: table;
}

.row::before {
  content: "";
  display: table;
  clear: both;
}

.row::after {
  content: "";
  display: table;
  clear: both;
}

.trow {
  display: table-row;
}

.cell {
  display: table-cell;
  vertical-align: middle;
}

@media (max-width:600px){
  .mobile{ display:block; width:100%!important }
  
  .container { padding: 0px 8px; }
  .container-small { padding: 0px 4px; }
  .padding-tiny { padding: 1px 2px; }
  .padding-small { padding: 2px 4px; }
  .padding-normal { padding: 4px 8px; }
  .padding-large { padding: 6px 12px; }
  .padding-xlarge { padding: 8px 16px;}

  .hide-small { display: none!important; }
} 

@media (min-width:601px) {
  .m1 { float:left; width: 8.33333%; } .m2 { float:left; width: 16.66666%; } .m3 { float:left; width:24.99999%; }
  .m4 { float:left; width: 33.33333%; } .m5 { float:left; width: 41.66666%; } .m6 { float:left; width: 49.99999%; }
  .m7 { float:left; width: 58.33333%; } .m8 { float:left; width: 66.66666%; } .m9 { float:left; width: 74.99999%; }
  .m10 { float:left; width: 83.33333%; } .m11 { float:left; width: 91.66666%; } .m12 { float:left; width: 99.99999%; }
  .half { width:49.99999% } .third{ width:33.33333% } .quarter{ width:24.99999% }
}

@media (max-width:992px) and (min-width:601px){
  .hide-medium { display: none!important; }
}

@media (min-width:993px) {
  .hide-large { display: none!important; }
  .l1 { float:left; width: 8.33333%; } .l2 { float:left; width: 16.66666%; } .l3 { float:left; width: 24.99999%; }
  .l4 { float:left; width: 33.33333%; } .l5 { float:left; width: 41.66666%; } .l6 {float:left; width: 49.99999%; }
  .l7 { float:left; width: 58.33333%; } .l8 { float:left; width: 66.66666%; } .l9 { float:left; width: 74.99999%; }
  .l10 { float:left; width:83.33333%; } .l11 { float:left; width: 91.66666%; } .l12 { float:left; width: 99.99999%; }
}
`
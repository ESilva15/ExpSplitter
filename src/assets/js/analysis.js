(function() {
    function loadScript(src, callback) {
      if (window.Chart && src.includes('chart.js')) {
        callback();
        return;
      }
      if (document.querySelector(`script[src="${src}"]`)) {
        if (window.Chart) {
          callback();
        } else {
          const existingScript = document.querySelector(`script[src="${src}"]`);
          existingScript.addEventListener('load', callback);
        }
        return;
      }

      const script = document.createElement('script');
      script.src = src;
      script.onload = callback;
      document.head.appendChild(script);
    }

    // Load Chart.js first, then the plugin, then render charts
    loadScript('/assets/js/chart.js', function() {
      loadScript('/assets/js/chartjs-plugin-datalabels.js', function() {
        renderAnalysisCharts();
      });
    });

    function renderAnalysisCharts() {
      // Read the global data injected by the HTML template
      const monthlyBalanceData = window.analysisBalanceData || [];
      const mTotalData = window.analysisTotalData || [];

      // --- Monthly Balance Chart ---
      const mBalanceChart = document.getElementById('monthlyBalance');
      const labels = monthlyBalanceData.map(entry =>
        new Date(2000, entry.month, 1).toLocaleString('en', {month: 'short'})
      );
      const values = monthlyBalanceData.map(entry => Number(entry.total));

      new Chart(mBalanceChart, {
        type: 'bar',
        data: {
          labels: labels,
          datasets: [{
            label: 'Monthly Balance',
            data: values,
            backgroundColor: (ctx) => {
              const value = ctx.raw;
              return value < 0 ? 'rgb(231 144 234 / 50%)' : 'rgb(67 149 84 / 50%)';
            },
            borderColor: (ctx) => {
              const value = ctx.raw;
              return value < 0 ? 'rgb(231 144 234 / 100%)' : 'rgb(67 149 84 / 100%)';
            },
            borderWidth: 1
          }]
        },
        options: {
          plugins: {
            datalabels: {
              anchor: "end",
              align: "end",
              formatter: (value) => value,
              color: (context) => {
                const value = context.dataset.data[context.dataIndex];
                return value < 0 ? 'rgb(231 144 234 / 100%)' : 'rgb(67 149 84 / 100%)';
              },
              font: { weight: 'bold' }
            }
          },
          scales: {
            y: { beginAtZero: true }
          }
        },
        plugins: [ChartDataLabels]
      });

      // --- Total Chart ---
      const mTotalChart = document.getElementById('monthlyTotal');
      const mTotalLabels = mTotalData.map(entry =>
        new Date(2000, entry.month, 1).toLocaleString('en', {month: 'short'})
      );
      const mTotalValues = mTotalData.map(entry => Number(entry.total));

      new Chart(mTotalChart, {
        type: 'line',
        data: {
          labels: mTotalLabels,
          datasets: [{
            label: 'Monthly Total',
            data: mTotalValues,
            backgroundColor: (ctx) => {
              const value = ctx.raw;
              return value < 0 ? 'rgb(231 144 234 / 50%)' : 'rgb(67 149 84 / 50%)';
            },
            borderColor: (ctx) => {
              const value = ctx.raw;
              return value < 0 ? 'rgb(231 144 234 / 100%)' : 'rgb(67 149 84 / 100%)';
            },
            borderWidth: 1
          }]
        },
        options: {
          plugins: {
            datalabels: {
              anchor: "end",
              align: "end",
              formatter: (value) => value,
              color: (context) => {
                const value = context.dataset.data[context.dataIndex];
                return value < 0 ? 'rgb(231 144 234 / 100%)' : 'rgb(67 149 84 / 100%)';
              },
              font: { weight: 'bold' }
            }
          },
          scales: {
            y: { beginAtZero: true }
          }
        },
        plugins: [ChartDataLabels]
      });
    }
  })();

#!/usr/bin/env python3
"""Generate a CLI command tree diagram for the CVE Utils project."""

import matplotlib
import os
matplotlib.use('Agg')
import matplotlib.pyplot as plt
from matplotlib.patches import FancyBboxPatch
import numpy as np


def draw_cli_tree():
    fig, ax = plt.subplots(1, 1, figsize=(26, 18))
    ax.set_xlim(-13, 13)
    ax.set_ylim(-10, 6)
    ax.set_aspect('equal')
    ax.axis('off')
    fig.patch.set_facecolor('#fafbfc')

    # Color palette
    COLORS = {
        'root':      '#2d3436',
        'format':    '#e94560',
        'validate':  '#ff6b6b',
        'extract':   '#4ecdc4',
        'compare':   '#45b7d1',
        'filter':    '#96ceb4',
        'generate':  '#f9ca24',
        'set':       '#a29bfe',
        'range':     '#fab1a0',
        'stats':     '#00cec9',
        'other':     '#6c5ce7',
    }

    def draw_box(x, y, w, h, label, sublabel, facecolor, edgecolor,
                 fontsize=11, subfontsize=8, text_color='white', sub_color='#b2bec3',
                 alpha=0.95, pad=0.15):
        box = FancyBboxPatch((x - w/2, y - h/2), w, h,
                              boxstyle=f"round,pad={pad}", facecolor=facecolor,
                              edgecolor=edgecolor, linewidth=2, zorder=5, alpha=alpha)
        ax.add_patch(box)
        if sublabel:
            ax.text(x, y + 0.12, label, fontsize=fontsize, fontweight='bold',
                    ha='center', va='center', zorder=6, color=text_color,
                    fontfamily='sans-serif')
            ax.text(x, y - 0.22, sublabel, fontsize=subfontsize,
                    ha='center', va='center', zorder=6, style='italic',
                    color=sub_color, fontfamily='sans-serif')
        else:
            ax.text(x, y, label, fontsize=fontsize, fontweight='bold',
                    ha='center', va='center', zorder=6, color=text_color,
                    fontfamily='monospace')

    def draw_leaf(x, y, label, edge_color, fontsize=8.5):
        text_w = max(len(label) * 0.09 + 0.2, 1.0)
        box = FancyBboxPatch((x - text_w/2, y - 0.22), text_w, 0.44,
                              boxstyle="round,pad=0.08", facecolor='white',
                              edgecolor=edge_color, linewidth=1.5, zorder=5, alpha=0.92)
        ax.add_patch(box)
        ax.text(x, y, label, fontsize=fontsize, color='#2d3436',
                ha='center', va='center', zorder=6, fontfamily='monospace',
                fontweight='bold')

    def draw_branch(x1, y1, x2, y2, color='#b2bec3', lw=2):
        t_vals = np.linspace(0, 1, 40)
        cx = (x1 + x2) / 2 + 0.3
        cy = (y1 + y2) / 2
        curve_x = (1-t_vals)**2 * x1 + 2*(1-t_vals)*t_vals * cx + t_vals**2 * x2
        curve_y = (1-t_vals)**2 * y1 + 2*(1-t_vals)*t_vals * cy + t_vals**2 * y2
        ax.plot(curve_x, curve_y, color=color, linewidth=lw, alpha=0.5, zorder=1,
                solid_capstyle='round')

    def draw_line(x1, y1, x2, y2, color='#b2bec3', lw=1.5):
        ax.plot([x1, x2], [y1, y2], color=color, linewidth=lw, alpha=0.4, zorder=1,
                solid_capstyle='round')

    # ===== Title =====
    ax.text(0, 5.3, 'CLI Command Tree', fontsize=28, fontweight='bold',
            color='#2d3436', ha='center', va='center', fontfamily='sans-serif')
    ax.text(0, 4.7, 'cve — Comprehensive CLI tool for CVE identifier processing',
            fontsize=13, color='#636e72', ha='center', va='center', style='italic')

    # ===== Root node =====
    root_x, root_y = 0, 3.5
    draw_box(root_x, root_y, 2.8, 0.85, 'cve', 'root command', COLORS['root'], '#1a1a2e',
             fontsize=16, pad=0.2)

    # ===== Row 1: Top-level commands with subcommands =====
    # --- extract (has subcmds) ---
    ex_x, ex_y = -10, 1.2
    draw_box(ex_x, ex_y, 2.4, 0.7, 'extract', 'Extract from text', COLORS['extract'], COLORS['extract'], fontsize=10)
    draw_branch(root_x, root_y - 0.425, ex_x, ex_y + 0.35, color=COLORS['extract'])

    ex_children = ['first', 'last', 'seq', 'split', 'year']
    for i, ch in enumerate(ex_children):
        cx = -12 + i * 1.2
        cy = -0.5
        draw_leaf(cx, cy, ch, COLORS['extract'])
        draw_line(ex_x, ex_y - 0.35, cx, cy + 0.22, color=COLORS['extract'])

    # --- filter (has subcmds) ---
    fi_x, fi_y = -4.5, 1.2
    draw_box(fi_x, fi_y, 2.4, 0.7, 'filter', 'Filter & group', COLORS['filter'], COLORS['filter'], fontsize=10)
    draw_branch(root_x, root_y - 0.425, fi_x, fi_y + 0.35, color=COLORS['filter'])

    fi_children = ['by-year', 'by-year-range', 'recent', 'dedup', 'group-by-year']
    for i, ch in enumerate(fi_children):
        cx = -7.5 + i * 1.5
        cy = -0.5
        draw_leaf(cx, cy, ch, COLORS['filter'])
        draw_line(fi_x, fi_y - 0.35, cx, cy + 0.22, color=COLORS['filter'])

    # --- generate (has subcmds) ---
    ge_x, ge_y = 3.5, 1.2
    draw_box(ge_x, ge_y, 2.4, 0.7, 'generate', 'Generate CVE IDs', COLORS['generate'], COLORS['generate'], fontsize=10)
    draw_branch(root_x, root_y - 0.425, ge_x, ge_y + 0.35, color=COLORS['generate'])

    ge_children = ['cve', 'fake']
    for i, ch in enumerate(ge_children):
        cx = 2.5 + i * 2.0
        cy = -0.5
        draw_leaf(cx, cy, ch, COLORS['generate'])
        draw_line(ge_x, ge_y - 0.35, cx, cy + 0.22, color=COLORS['generate'])

    # ===== Row 2: Standalone commands (no subcmds) =====
    standalone_cmds = [
        (-10.5, -2.5, 'format',       'Format CVE IDs',          COLORS['format'],  2.2),
        (-8.0,  -2.5, 'validate',     'Validate single CVE',     COLORS['validate'], 2.2),
        (-5.5,  -2.5, 'validate-batch','Batch validate CVEs',    COLORS['validate'], 2.6),
        (-2.8,  -2.5, 'filter-pattern','Pattern filter',         COLORS['filter'],   2.6),
        (-0.2,  -2.5, 'filter-valid',  'Filter valid CVEs',      COLORS['filter'],   2.4),
        (2.5,   -2.5, 'compare',      'Compare two CVEs',        COLORS['compare'],  2.2),
        (5.0,   -2.5, 'format-seq',   'Zero-pad seq number',     COLORS['format'],   2.4),
    ]

    for sx, sy, label, sublabel, color, w in standalone_cmds:
        draw_box(sx, sy, w, 0.65, label, sublabel, color, color, fontsize=9, subfontsize=7)
        draw_branch(root_x, root_y - 0.425, sx, sy + 0.325, color=color, lw=1.5)

    # ===== Row 3: Set ops, Range, Stats =====
    row3_cmds = [
        (-10.5, -5.0, 'intersect',     'Set intersection',    COLORS['set'],   2.2),
        (-8.0,  -5.0, 'union',         'Set union',           COLORS['set'],   2.0),
        (-5.5,  -5.0, 'diff',          'Set difference',      COLORS['set'],   2.0),
        (-3.0,  -5.0, 'parse-range',   'Parse CVE range',     COLORS['range'], 2.4),
        (-0.5,  -5.0, 'is-consecutive','Check consecutive',   COLORS['range'], 2.6),
        (2.0,   -5.0, 'count-by-year', 'Count per year',      COLORS['stats'], 2.4),
        (4.5,   -5.0, 'year-range',    'Year span',           COLORS['stats'], 2.2),
        (7.0,   -5.0, 'seq-range',     'Seq number range',    COLORS['stats'], 2.2),
    ]

    for sx, sy, label, sublabel, color, w in row3_cmds:
        draw_box(sx, sy, w, 0.65, label, sublabel, color, color, fontsize=9, subfontsize=7)
        draw_branch(root_x, root_y - 0.425, sx, sy + 0.325, color=color, lw=1.5)

    # ===== Row 4: Other commands =====
    other_cmds = [
        (-10.0, -7.5, 'version',    'Print version',     COLORS['other'], 2.0),
        (-7.0,  -7.5, 'completion', 'Shell completion',  COLORS['other'], 2.2),
    ]

    for sx, sy, label, sublabel, color, w in other_cmds:
        draw_box(sx, sy, w, 0.55, label, sublabel, color, color, fontsize=8, subfontsize=6)
        draw_branch(root_x, root_y - 0.425, sx, sy + 0.275, color=color, lw=1.2)

    # ===== Legend =====
    legend_items = [
        ('Format & Validate', COLORS['format']),
        ('Extract', COLORS['extract']),
        ('Compare', COLORS['compare']),
        ('Filter & Group', COLORS['filter']),
        ('Generate', COLORS['generate']),
        ('Set Operations', COLORS['set']),
        ('Range & Pattern', COLORS['range']),
        ('Statistics', COLORS['stats']),
        ('Other', COLORS['other']),
    ]
    legend_y = -9.2
    legend_spacing = 2.6
    for i, (label, color) in enumerate(legend_items):
        lx = -11.5 + i * legend_spacing
        box = FancyBboxPatch((lx - 0.2, legend_y - 0.18), 0.36, 0.36,
                              boxstyle="round,pad=0.06", facecolor=color,
                              edgecolor='white', linewidth=1, zorder=5, alpha=0.9)
        ax.add_patch(box)
        ax.text(lx + 0.35, legend_y, label, fontsize=8, color='#2d3436',
                ha='left', va='center', fontfamily='sans-serif')

    # ===== Usage hint =====
    ax.text(0, -9.8, 'Usage: cve <command> [flags]    |    Install: go install github.com/scagogogo/cve-skills/cmd/cve@latest',
            fontsize=10, color='#636e72', ha='center', va='center',
            fontfamily='monospace', style='italic',
            bbox=dict(boxstyle='round,pad=0.4', facecolor='white',
                      edgecolor='#b2bec3', linewidth=1, alpha=0.8))

    plt.tight_layout(pad=0.5)
    plt.savefig(os.path.join(os.path.dirname(os.path.abspath(__file__)), '..', 'docs', 'images', 'cli-tree.png'),
                dpi=150, bbox_inches='tight', facecolor='#fafbfc')
    print("CLI tree diagram saved!")


if __name__ == '__main__':
    draw_cli_tree()
